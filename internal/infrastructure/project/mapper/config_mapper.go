// Package mapper provides functions to map between DTOs and domain models.
package mapper

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/dto"
)

func ToDomain(configDto dto.FileConfig) *vos.Config {
	if configDto.Runtime.Volumes.ProjectMountPath == "" {
		configDto.Runtime.Volumes.ProjectMountPath = vos.DefaultProjectMountPath
	}
	if configDto.Runtime.Volumes.StateMountPath == "" {
		configDto.Runtime.Volumes.StateMountPath = vos.DefaultStateMountPath
	}
	if configDto.State.Backend == "" {
		configDto.State.Backend = vos.DefaultStateBackend
	}

	return &vos.Config{
		Project: vos.Project{
			ID:           configDto.Project.ID,
			Name:         configDto.Project.Name,
			Version:      configDto.Project.Version,
			Team:         configDto.Project.Team,
			Description:  configDto.Project.Description,
			Organization: configDto.Project.Organization,
		},
		Template: vos.Template{
			URL: configDto.Template.URL,
			Ref: configDto.Template.Ref,
		},
		Technology: vos.Technology{
			Stack:          configDto.Technology.Stack,
			Infrastructure: configDto.Technology.Infrastructure,
		},
		Runtime: vos.Runtime{
			Image: vos.Image{
				Source: configDto.Runtime.Image.Source,
				Tag:    configDto.Runtime.Image.Tag,
			},
			Volumes: vos.Volumes{
				ProjectMountPath: configDto.Runtime.Volumes.ProjectMountPath,
				StateMountPath:   configDto.Runtime.Volumes.StateMountPath,
			},
		},
		State: vos.State{
			Backend: configDto.State.Backend,
			URL:     configDto.State.URL,
		},
	}
}
