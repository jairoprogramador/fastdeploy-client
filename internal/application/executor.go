package application

import (
	"context"
	"errors"

	appPor "github.com/jairoprogramador/fastdeploy/internal/application/ports"
	docAgg "github.com/jairoprogramador/fastdeploy/internal/domain/docker/aggregates"
	docPor "github.com/jairoprogramador/fastdeploy/internal/domain/docker/ports"
	proPor "github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
)

const MessageProjectNotInitialized = "project not initialized. Please run 'fd init' first"

type ExecutorService struct {
	projectRepository  proPor.ProjectRepository
	dockerService      docPor.DockerService
	variableResolver   appPor.VarsResolver
	coreVersion        appPor.CoreVersion
	isTerminal         bool
	hostProjectPath    string
	hostFastdeployPath string
}

func NewExecutorService(
	isTerminal bool,
	hostProjectPath string,
	hostFastdeployPath string,
	projectRepository proPor.ProjectRepository,
	dockerService docPor.DockerService,
	variableResolver appPor.VarsResolver,
	coreVersion appPor.CoreVersion,
) *ExecutorService {
	return &ExecutorService{
		isTerminal:         isTerminal,
		hostProjectPath:    hostProjectPath,
		hostFastdeployPath: hostFastdeployPath,
		projectRepository:  projectRepository,
		dockerService:      dockerService,
		variableResolver:   variableResolver,
		coreVersion:        coreVersion,
	}
}

func (s *ExecutorService) Run(ctx context.Context, command, environment string, withTty bool) error {
	// 1. Validación de pre-condiciones
	exists, err := s.projectRepository.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(MessageProjectNotInitialized)
	}

	if err := s.dockerService.Check(ctx); err != nil {
		return err
	}

	// 2. Cargar el estado del dominio
	project, err := s.projectRepository.Load()
	if err != nil {
		return err
	}

	latestCoreVersion, _ := s.coreVersion.GetLatestVersion()

	// 3. Delegar la lógica de negocio al dominio 'docker'
	execution, err := docAgg.NewExecution(
		project,
		command,
		environment,
		s.hostProjectPath,
		s.hostFastdeployPath,
		s.isTerminal,
		withTty,
		latestCoreVersion,
	)
	if err != nil {
		return err
	}

	// 4. Orquestar la ejecución
	if execution.NeedsBuild() {
		buildOpts := execution.BuildOptions()
		if err := s.dockerService.Build(ctx, *buildOpts); err != nil {
			return err
		}
	}

	// La resolución de variables de entorno se queda en la capa de aplicación,
	// ya que depende de variables internas que se generan en tiempo de ejecución (que ya no tenemos, como auth)
	// pero que podrían existir en el futuro.
	runOpts := execution.RunOptions()
	resolvedEnvVars := make(map[string]string)
	for name, value := range runOpts.EnvVars {
		resolvedEnvVars[name] = s.variableResolver.Resolve(value, make(map[string]string))
	}
	runOpts.EnvVars = resolvedEnvVars

	_, err = s.dockerService.Run(ctx, runOpts)
	return err
}
