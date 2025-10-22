package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
)

const MessageProjectAlreadyExists = "project already initialized, fdconfig.yaml exists"
var ErrProjectAlreadyExists = errors.New(MessageProjectAlreadyExists)

type InitService struct {
	projectRepository ports.ProjectRepository
	inputService      ports.UserInputService
	projectName           string
}

func NewInitService(projectName string, repository ports.ProjectRepository, inputSvc ports.UserInputService) *InitService {
	return &InitService{
		projectRepository: repository,
		inputService:      inputSvc,
		projectName:       projectName,
	}
}

func (s *InitService) InitializeProject(ctx context.Context, interactive bool) error {
	exists, err := s.projectRepository.Exists()
	if err != nil {
		return fmt.Errorf("could not check if project exists: %w", err)
	}
	if exists {
		return ErrProjectAlreadyExists
	}

	var cfg *vos.Config
	if interactive {
		cfg, err = s.gatherConfigFromUser(ctx)
	} else {
		cfg, err = s.gatherDefaultConfig()
	}

	if err != nil {
		return err
	}

	return s.projectRepository.Save(cfg)
}

func (s *InitService) gatherConfigFromUser(ctx context.Context) (*vos.Config, error) {
	cfg, err := s.gatherDefaultConfig()
	if err != nil {
		return nil, err
	}

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

func (s *InitService) gatherDefaultConfig() (*vos.Config, error) {
	return &vos.Config{
		Project: vos.Project{
			Name:         s.projectName,
			Version:      vos.DefaultProjectVersion,
			Team:         vos.DefaultProjectTeam,
			Description:  vos.DefaultProjectDescription,
			Organization: vos.DefaultProjectOrganization,
		},
		Template: vos.Template{
			URL: vos.DefaultUrl,
			Ref: vos.DefaultRef,
		},
		Technology: vos.Technology{
			Stack:          vos.DefaultStack,
			Infrastructure: vos.DefaultInfrastructure,
		},
		Runtime: vos.Runtime{
			Image: vos.Image{
				Source: vos.DefaultImageSource,
				Tag:    vos.DefaultImageTag,
			},
			Volumes: vos.Volumes{
				ProjectMountPath: vos.DefaultProjectMountPath,
				StateMountPath:   vos.DefaultStateMountPath,
			},
		},
		State: vos.State{
			Backend: vos.DefaultStateBackend,
			URL:     vos.DefaultStateURL,
		},
	}, nil
}
