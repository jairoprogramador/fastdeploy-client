package application

import (
	"context"
	"fmt"

	proPor "github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	proVos "github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"

	appPor "github.com/jairoprogramador/fastdeploy/internal/application/ports"
)

const MessageProjectAlreadyExists = "project already initialized, fdconfig.yaml exists"

type InitService struct {
	projectRepository proPor.ProjectRepository
	inputService      proPor.UserInputService
	projectName       string
	logMessage        appPor.LogMessage
}

func NewInitService(
	projectName string,
	repository proPor.ProjectRepository,
	inputSvc proPor.UserInputService,
	logMessage appPor.LogMessage) *InitService {
	return &InitService{
		projectRepository: repository,
		inputService:      inputSvc,
		projectName:       projectName,
		logMessage:        logMessage,
	}
}

func (s *InitService) InitializeProject(ctx context.Context, interactive bool) error {
	s.logMessage.Info("initializing project...")
	exists, err := s.projectRepository.Exists()
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return err
	}
	if exists {
		s.logMessage.Info(MessageProjectAlreadyExists)
		return nil
	}

	var cfg *proVos.Config
	if interactive {
		cfg, err = s.gatherConfigFromUser(ctx)
		if err != nil {
			s.logMessage.Error(fmt.Sprintf("%v", err))
			return err
		}
	} else {
		cfg = s.gatherDefaultConfig()
	}

	err = s.projectRepository.Save(cfg)
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return err
	}

	s.logMessage.Success("project initialized successfully")
	return nil
}

func (s *InitService) gatherConfigFromUser(ctx context.Context) (*proVos.Config, error) {
	cfg := s.gatherDefaultConfig()
	
    var err error

	cfg.Project.Name, err = s.inputService.Ask(ctx, "Project Name", cfg.Project.Name)
	if err != nil {
		return nil, err
	}
	cfg.Project.Version, err = s.inputService.Ask(ctx, "Project Version", cfg.Project.Version)
	if err != nil {
		return nil, err
	}
	cfg.Project.Team, err = s.inputService.Ask(ctx, "Project Team", cfg.Project.Team)
	if err != nil {
		return nil, err
	}
	cfg.Project.Organization, err = s.inputService.Ask(ctx, "Project Organization", cfg.Project.Organization)
	if err != nil {
		return nil, err
	}
	cfg.Template.URL, err = s.inputService.Ask(ctx, "Template URL", cfg.Template.URL)
	if err != nil {
		return nil, err
	}
	cfg.Technology.Stack, err = s.inputService.Ask(ctx, "Technology Stack", cfg.Technology.Stack)
	if err != nil {
		return nil, err
	}
	cfg.Technology.Infrastructure, err = s.inputService.Ask(ctx, "Technology Infrastructure", cfg.Technology.Infrastructure)
	if err != nil {
		return nil, err
	}
	cfg.Runtime.Image.Source, err = s.inputService.Ask(ctx, "Runtime Image Source", cfg.Runtime.Image.Source)
	if err != nil {
		return nil, err
	}
	cfg.Runtime.Image.Tag, err = s.inputService.Ask(ctx, "Runtime Image Tag", cfg.Runtime.Image.Tag)
	if err != nil {
		return nil, err
	}
	cfg.State.Backend, err = s.inputService.Ask(ctx, "State Backend", cfg.State.Backend)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (s *InitService) gatherDefaultConfig() *proVos.Config {
	return &proVos.Config{
		Project: proVos.Project{
			Name:         s.projectName,
			Version:      proVos.DefaultProjectVersion,
			Team:         proVos.DefaultProjectTeam,
			Description:  proVos.DefaultProjectDescription,
			Organization: proVos.DefaultProjectOrganization,
		},
		Template: proVos.Template{
			URL: proVos.DefaultUrl,
			Ref: proVos.DefaultRef,
		},
		Technology: proVos.Technology{
			Stack:          proVos.DefaultStack,
			Infrastructure: proVos.DefaultInfrastructure,
		},
		Runtime: proVos.Runtime{
			Image: proVos.Image{
				Source: proVos.DefaultImageSource,
				Tag:    proVos.DefaultImageTag,
			},
			Volumes: proVos.Volumes{
				ProjectMountPath: proVos.DefaultProjectMountPath,
				StateMountPath:   proVos.DefaultStateMountPath,
			},
		},
		State: proVos.State{
			Backend: proVos.DefaultStateBackend,
			URL:     proVos.DefaultStateURL,
		},
	}
}
