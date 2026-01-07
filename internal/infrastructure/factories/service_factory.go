package factories

import (
	"os"
	"path/filepath"

	app "github.com/jairoprogramador/fastdeploy/internal/application"
	dockerDomain "github.com/jairoprogramador/fastdeploy/internal/domain/docker/services"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	//"github.com/jairoprogramador/fastdeploy/internal/infrastructure/logger"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/docker"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project"
)

type ServiceFactory interface {
	BuildExecutor() (*app.ExecutorService, error)
	BuildInitialize() (*app.InitializeService, error)
	//BuildPresenter() app.Presenter
}

type serviceFactory struct{}

func NewServiceFactory() ServiceFactory {
	return &serviceFactory{}
}

/* func (f *serviceFactory) BuildPresenter() app.Presenter {
	return logger.NewConsolePresenter()
} */

func (f *serviceFactory) BuildInitialize() (*app.InitializeService, error) {
	projectPath, err := f.getProjectPath()
	if err != nil {
		return nil, err
	}
	projectRepository, err := f.getProjectRepository(projectPath)
	if err != nil {
		return nil, err
	}

	inputService := project.NewSurveyUserInputService()
	versionService := project.NewHttpVersion()

	return app.NewInitializeService(
		filepath.Base(projectPath), projectRepository, inputService, versionService), nil
}

func (f *serviceFactory) BuildExecutor() (*app.ExecutorService, error) {
	projectPath, err := f.getProjectPath()
	if err != nil {
		return nil, err
	}
	projectRepository, err := f.getProjectRepository(projectPath)
	if err != nil {
		return nil, err
	}

	cmdExecutor := docker.NewShellExecutor()
	imageService := dockerDomain.NewImageBuilder()
	containerService := dockerDomain.NewContainerBuilder()

	return app.NewExecutorService(
		projectRepository, cmdExecutor, imageService, containerService), nil
}

func (f *serviceFactory) getProjectRepository(projectPath string) (ports.ProjectRepository, error) {
	projectRepository := project.NewYAMLProjectRepository(projectPath)
	return projectRepository, nil
}

func (f *serviceFactory) getProjectPath() (string, error) {
	return os.Getwd()
}