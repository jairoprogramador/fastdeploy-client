// Package mapper provides functions to map between DTOs and domain models.
package mapper

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/dto"
)

func ToDomain(fc dto.FileConfig) *vos.Config {
	if fc.Runtime.Volumes.ProjectMountPath == "" {
		fc.Runtime.Volumes.ProjectMountPath = vos.DefaultProjectMountPath
	}
	if fc.Runtime.Volumes.StateMountPath == "" {
		fc.Runtime.Volumes.StateMountPath = vos.DefaultStateMountPath
	}
	if fc.State.Backend == "" {
		fc.State.Backend = vos.DefaultStateBackend
	}

	return &vos.Config{
		Project: vos.Project{
			Name:         fc.Project.Name,
			Version:      fc.Project.Version,
			Team:         fc.Project.Team,
			Description:  fc.Project.Description,
			Organization: fc.Project.Organization,
		},
		Template: vos.Template{
			URL: fc.Template.URL,
			Ref: fc.Template.Ref,
		},
		Technology: vos.Technology{
			Stack:          fc.Technology.Stack,
			Infrastructure: fc.Technology.Infrastructure,
		},
		Runtime: vos.Runtime{
			Image: vos.Image{
				Source: fc.Runtime.Image.Source,
				Tag:    fc.Runtime.Image.Tag,
			},
			Volumes: vos.Volumes{
				ProjectMountPath: fc.Runtime.Volumes.ProjectMountPath,
				StateMountPath:   fc.Runtime.Volumes.StateMountPath,
			},
		},
		State: vos.State{
			Backend: fc.State.Backend,
			URL:     fc.State.URL,
		},
	}
}
