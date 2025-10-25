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
)

const MessageProjectNotInitialized = "project not initialized. Please run 'fd init' first"

type OrderService struct {
	isTerminal        bool
	workDir           string
	fastdeployHome    string
	projectRepository proPor.ProjectRepository
	dockerService     docPor.DockerService
	logMessage        appPor.LogMessage
}

func NewOrderService(
	isTerminal bool,
	workDir string,
	fastdeployHome string,
	projectRepository proPor.ProjectRepository,
	dockerService docPor.DockerService,
	logMessage appPor.LogMessage,
) *OrderService {
	return &OrderService{
		isTerminal:        isTerminal,
		workDir:           workDir,
		fastdeployHome:    fastdeployHome,
		projectRepository: projectRepository,
		dockerService:     dockerService,
		logMessage:        logMessage,
	}
}

func (s *OrderService) ExecuteOrder(ctx context.Context, order, env string, withTty bool) error {
	s.logMessage.Info(fmt.Sprintf("executing Order: %s", order))

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

	runOpts := s.prepareRunOptions(fileConfig, localImage, s.workDir, order, env, withTty)
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

func (s *OrderService) prepareRunOptions(fileConfig *proVos.Config, image docVos.Image, workDir, order, env string, withTty bool) docVos.RunOptions {

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

	return docVos.RunOptions{
		Image:        image,
		Volumes:      volumes,
		EnvVars:      envVars,
		Command:      strings.TrimSpace(fmt.Sprintf("%s %s", order, env)),
		Interactive:  interactive,
		AllocateTTY:  allocateTty,
		RemoveOnExit: true,
	}
}
