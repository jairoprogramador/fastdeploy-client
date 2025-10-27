package application

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	appPor "github.com/jairoprogramador/fastdeploy/internal/application/ports"
	docPor "github.com/jairoprogramador/fastdeploy/internal/domain/docker/ports"
	docVos "github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	proPor "github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	proVos "github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	fdplug "github.com/jairoprogramador/fastdeploy/internal/fdplugin"
)

const MessageProjectNotInitialized = "project not initialized. Please run 'fd init' first"

type OrderService struct {
	isTerminal        bool
	workDir           string
	fastdeployHome    string
	projectRepository proPor.ProjectRepository
	dockerService     docPor.DockerService
	authService       appPor.AuthService
	variableResolver  appPor.VarsResolver
	logMessage        appPor.LogMessage
}

func NewOrderService(
	isTerminal bool,
	workDir string,
	fastdeployHome string,
	projectRepository proPor.ProjectRepository,
	dockerService docPor.DockerService,
	authService appPor.AuthService,
	variableResolver appPor.VarsResolver,
	logMessage appPor.LogMessage,
) *OrderService {
	return &OrderService{
		isTerminal:        isTerminal,
		workDir:           workDir,
		fastdeployHome:    fastdeployHome,
		projectRepository: projectRepository,
		dockerService:     dockerService,
		authService:       authService,
		variableResolver:  variableResolver,
		logMessage:        logMessage,
	}
}

func (s *OrderService) ExecuteOrder(ctx context.Context, order, env string, withTty bool) error {
	s.logMessage.Info(fmt.Sprintf("executing order: %s", order))

	exists, err := s.projectRepository.Exists()
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return err
	}
	if !exists {
		s.logMessage.Info(MessageProjectNotInitialized)
		return nil
	}

	if err := s.dockerService.Check(ctx); err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return err
	}

	fileConfig, err := s.projectRepository.Load()
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return err
	}

	internalVars := make(map[string]string)

	if fileConfig.Auth.Plugin != "" {
		s.logMessage.Info("Authenticating...")
		
		resolvedParams := &fdplug.AuthConfig{
			ClientId:  s.variableResolver.Resolve(fileConfig.Auth.Params.ClientID, internalVars),
			ClientSecret: s.variableResolver.Resolve(fileConfig.Auth.Params.ClientSecret, internalVars),
			GrantType: fdplug.AuthGrantType(fdplug.AuthGrantType_value[fileConfig.Auth.Params.GrantType]),
			Extra:     make(map[string]string),
			Scope:     fileConfig.Auth.Params.Scope,
		}
		for key, val := range fileConfig.Auth.Params.Extra {
			resolvedParams.Extra[key] = s.variableResolver.Resolve(val, internalVars)
		}

		authenticateRequest := &fdplug.AuthenticateRequest{
			Config: resolvedParams,
		}

		authResp, err := s.authService.Authenticate(ctx, fileConfig.Auth.Plugin, authenticateRequest)
		if err != nil {
			s.logMessage.Error(fmt.Sprintf("Authentication failed: %v", err))
			return err
		}

		tokenVarName := strings.ToUpper(fileConfig.Auth.Plugin) + "_ACCESS_TOKEN"
		internalVars[tokenVarName] = authResp.Token.AccessToken
		s.logMessage.Success("Authentication successful")
	}

	var localImage docVos.Image
	if fileConfig.Runtime.Image.Source != "" {
		localImage = docVos.Image{
			Name: fileConfig.Runtime.Image.Source,
			Tag:  fileConfig.Runtime.Image.Tag}
	} else {
		buildOpts, localImageBuilt := s.prepareBuildOptions(fileConfig)
		if err := s.dockerService.Build(ctx, buildOpts); err != nil {
			return err
		}
		localImage = localImageBuilt
	}

	runOpts := s.prepareRunOptions(fileConfig, localImage, s.workDir, order, env, withTty, internalVars)
	if err := s.dockerService.Run(ctx, runOpts); err != nil {
		return err
	}

	s.logMessage.Success("Order executed successfully")
	return nil
}

func (s *OrderService) prepareBuildOptions(fileConfig *proVos.Config) (docVos.BuildOptions, docVos.Image) {
	localImageName := fmt.Sprintf("%s-%s",
		fileConfig.Project.Team,
		fileConfig.Template.NameTemplate(),
	)
	localImage := docVos.Image{
		Name: localImageName,
		Tag:  fileConfig.Runtime.Image.Tag}

	buildArgs := make(map[string]string)
	if runtime.GOOS == "linux" {
		buildArgs["DEV_GID"] = "$(id -g)"
	}

	if fileConfig.Runtime.CoreVersion != "" {
		buildArgs["FASTDEPLOY_VERSION"] = fileConfig.Runtime.CoreVersion
	}

	return docVos.BuildOptions{
		Image:      localImage,
		Context:    ".",
		Dockerfile: "Dockerfile",
		Args:       buildArgs,
	}, localImage
}

func (s *OrderService) prepareRunOptions(
	fileConfig *proVos.Config,
	image docVos.Image,
	workDir,
	order,
	env string,
	withTty bool,
	internalVars map[string]string,
) docVos.RunOptions {

	volumesMap := make(map[string]string)

	for _, volume := range fileConfig.Runtime.Volumes {
		volumesMap[volume.Host] = volume.Container
	}

	projectContainerPath, okProjectContainerPath := volumesMap[proVos.ProjectPathKey]
	if !okProjectContainerPath {
		volumesMap[workDir] = proVos.DefaultContainerProjectPath
	} else {
		volumesMap[workDir] = projectContainerPath
		delete(volumesMap, proVos.ProjectPathKey)
	}

	envVars := make(map[string]string)

	for _, envVar := range fileConfig.Runtime.Env {
		envVars[envVar.Name] = s.variableResolver.Resolve(envVar.Value, internalVars)
	}

	if fileConfig.State.Backend == proVos.DefaultStateBackend {

		stateContainerPath, okStateContainerPath := volumesMap[proVos.StatePathKey]
		if !okStateContainerPath {
			volumesMap[s.fastdeployHome] = proVos.DefaultContainerStatePath
		} else {
			volumesMap[s.fastdeployHome] = stateContainerPath
			delete(volumesMap, proVos.StatePathKey)
		}

		envVars["FASTDEPLOY_HOME"] = volumesMap[s.fastdeployHome]
	}

	volumes := make([]docVos.Volume, 0, len(volumesMap))
	for hostPath, containerPath := range volumesMap {
		volumes = append(volumes, docVos.Volume{
			HostPath:      hostPath,
			ContainerPath: containerPath,
		})
	}

	allocateTty := withTty && s.isTerminal
	interactive := allocateTty

	groups := []string{}
	if runtime.GOOS == "linux" {
		groups = append(groups, "$(getent group docker | cut -d: -f3)")
	}


	return docVos.RunOptions{
		Image:        image,
		Volumes:      volumes,
		EnvVars:      envVars,
		Command:      strings.TrimSpace(fmt.Sprintf("%s %s", order, env)),
		Interactive:  interactive,
		AllocateTTY:  allocateTty,
		RemoveOnExit: true,
		Groups:		  groups,
	}
}
