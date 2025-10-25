// Package mapper provides functions to map between DTOs and domain models.
package mapper

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/dto"
)

func ToDomain(configDto dto.FileConfig) *vos.Config {

	if configDto.State.Backend == "" {
		configDto.State.Backend = vos.DefaultStateBackend
	}

	domainVolumes := make([]vos.Volume, 0, len(configDto.Runtime.Volumes))
	for _, dtoVol := range configDto.Runtime.Volumes {
		domainVolumes = append(domainVolumes, vos.Volume{
			Host:      dtoVol.Host,
			Container: dtoVol.Container,
		})
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
		Template: vos.NewTemplate(configDto.Template.URL, configDto.Template.Ref),
		Runtime: vos.Runtime{
			CoreVersion: configDto.Runtime.CoreVersion,
			Image: vos.Image{
				Source: configDto.Runtime.Image.Source,
				Tag:    configDto.Runtime.Image.Tag,
			},
			Volumes: domainVolumes,
		},
		State: vos.State{
			Backend: configDto.State.Backend,
			URL:     configDto.State.URL,
		},
	}
}

func ToDto(config *vos.Config) dto.FileConfig {

	dtoVolumes := make([]dto.VolumeDTO, 0, len(config.Runtime.Volumes))
	for _, vol := range config.Runtime.Volumes {
		dtoVolumes = append(dtoVolumes, dto.VolumeDTO {
			Host:      vol.Host,
			Container: vol.Container,
		})
	}
	dtoConfig := dto.FileConfig{
		Project: dto.ProjectDTO{
			ID:           config.Project.ID,
			Name:         config.Project.Name,
			Version:      config.Project.Version,
			Team:         config.Project.Team,
			Description:  config.Project.Description,
			Organization: config.Project.Organization,
		},
		Template: dto.TemplateDTO{
			URL: config.Template.URL(),
			Ref: config.Template.Ref(),
		},
		Runtime: dto.RuntimeDTO{
			CoreVersion: config.Runtime.CoreVersion,
			Image: dto.ImageDTO{
				Source: config.Runtime.Image.Source,
				Tag:    config.Runtime.Image.Tag,
			},
			Volumes: dtoVolumes,
		},
		State: dto.StateDTO{
			Backend: config.State.Backend,
			URL:     config.State.URL,
		},
	}
	return dtoConfig
}
