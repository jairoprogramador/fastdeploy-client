package application

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"

	dockerPorts "github.com/jairoprogramador/fastdeploy/internal/domain/docker/ports"
	dockerVOs "github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	projectPorts "github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	projectVOs "github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
)

type RunStepService struct {
	isTerminal        bool
	workDir           string
	fastdeployHome    string
	projectRepository projectPorts.ProjectRepository
	dockerService     dockerPorts.DockerService
}

func NewRunStepService(
	isTerminal bool,
	workDir string,
	fastdeployHome string,
	projectRepo projectPorts.ProjectRepository,
	dockerService dockerPorts.DockerService,
) *RunStepService {
	return &RunStepService{
		isTerminal:        isTerminal,
		workDir:           workDir,
		fastdeployHome:    fastdeployHome,
		projectRepository: projectRepo,
		dockerService:     dockerService,
	}
}

func (s *RunStepService) ExecuteStep(ctx context.Context, step, env string, withTty bool) error {
	log.Println("Executing step:", step)

	if err := s.dockerService.Check(ctx); err != nil {
		return err
	}

	fileConfig, err := s.projectRepository.Load()
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	var localImage dockerVOs.Image
	if fileConfig.Runtime.Image.Source != "" {
		localImage = dockerVOs.Image{
			Name: fileConfig.Runtime.Image.Source,
			Tag:  fileConfig.Runtime.Image.Tag}
	} else {
		buildOpts, localImageBuilt := s.prepareBuildOptions(fileConfig)
		if err := s.dockerService.Build(ctx, buildOpts); err != nil {
			return err
		}
		localImage = localImageBuilt
	}

	runOpts := s.prepareRunOptions(fileConfig, localImage, s.workDir, step, env, withTty)
	if err := s.dockerService.Run(ctx, runOpts); err != nil {
		return err
	}

	log.Println("✅ Step executed successfully")
	return nil
}

func (s *RunStepService) prepareBuildOptions(fileConfig *projectVOs.Config) (dockerVOs.BuildOptions, dockerVOs.Image) {
	localImageName := fmt.Sprintf("%s-%s",
		fileConfig.Project.Team,
		fileConfig.Technology.Stack,
	)
	localImage := dockerVOs.Image{
		Name: localImageName,
		Tag:  fileConfig.Runtime.Image.Tag}

	buildArgs := make(map[string]string)
	if runtime.GOOS == "linux" {
		buildArgs["DEV_GID"] = "$(id -g)"
	}

	return dockerVOs.BuildOptions{
		Image:      localImage,
		Context:    ".",
		Dockerfile: "Dockerfile",
		Args:       buildArgs,
	}, localImage
}

func (s *RunStepService) prepareRunOptions(cfg *projectVOs.Config, image dockerVOs.Image, workDir, step, env string, withTty bool) dockerVOs.RunOptions {

	volumes := []dockerVOs.Volume{
		{HostPath: workDir, ContainerPath: cfg.Runtime.Volumes.ProjectMountPath},
	}

	envVars := make(map[string]string)

	if cfg.State.Backend == projectVOs.DefaultStateBackend {
		volumes = append(volumes, dockerVOs.Volume{
			HostPath:      s.fastdeployHome,
			ContainerPath: cfg.Runtime.Volumes.StateMountPath,
		})
		envVars["FASTDEPLOY_HOME"] = cfg.Runtime.Volumes.StateMountPath
	}

	allocateTty := withTty && s.isTerminal
	interactive := allocateTty

	return dockerVOs.RunOptions{
		Image:        image,
		Volumes:      volumes,
		EnvVars:      envVars,
		Command:      strings.TrimSpace(fmt.Sprintf("%s %s", step, env)),
		Interactive:  interactive,
		AllocateTTY:  allocateTty,
		RemoveOnExit: true,
	}
}
