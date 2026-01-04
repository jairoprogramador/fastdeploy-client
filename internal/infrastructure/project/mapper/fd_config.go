package mapper

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/aggregates"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/dto"
)

func ToDomainProject(configDto dto.ProjectDTO) (vos.ProjectID, vos.ProjectData, error) {
	id, err := vos.NewProjectID(configDto.ID)
	if err != nil {
		return vos.ProjectID{}, vos.ProjectData{}, err
	}
	data, err := vos.NewProjectData(
		configDto.Name,
		configDto.Organization,
		configDto.Team,
		configDto.Description)

	if err != nil {
		return vos.ProjectID{}, vos.ProjectData{}, err
	}
	return id, data, nil
}

func ToDomainAuth(configDto dto.AuthDTO) (vos.Auth, error) {
	extra := make(map[string]string)
	for _, item := range configDto.Params.Extra {
		for key, value := range item {
			extra[key] = value
		}
	}
	params, err := vos.NewAuthParams(
		configDto.Params.ClientID,
		configDto.Params.GrantType,
		configDto.Params.ClientSecret,
		configDto.Params.Scope,
		extra)

	if err != nil {
		return vos.Auth{}, err
	}
	return vos.NewAuth(configDto.Plugin, params), nil
}

func ToDomainRuntime(configDto dto.RuntimeDTO) (vos.Runtime, error) {
	container, err := vos.NewContainer(
		configDto.Container.Image,
		configDto.Container.Tag,
		configDto.Container.CoreVersion)
	if err != nil {
		return vos.Runtime{}, err
	}

	volumes := make([]vos.Volume, 0, len(configDto.Volumes))
	for _, dtoVol := range configDto.Volumes {
		volume, err := vos.NewVolume(dtoVol.Host, dtoVol.Container)
		if err != nil {
			return vos.Runtime{}, err
		}
		volumes = append(volumes, volume)
	}

	envVars := make([]vos.EnvVar, 0, len(configDto.Env))
	for _, dtoEnv := range configDto.Env {
		envVar, err := vos.NewEnvVar(dtoEnv.Name, dtoEnv.Value)
		if err != nil {
			return vos.Runtime{}, err
		}
		envVars = append(envVars, envVar)
	}

	runtime := vos.NewRuntime(container, volumes, envVars)
	return runtime, nil
}

func ToDomain(configDto dto.FDConfigDTO) (*aggregates.Project, error) {

	id, data, err := ToDomainProject(configDto.Project)
	if err != nil {
		return nil, err
	}

	template, err := vos.NewTemplate(configDto.Template.URL, configDto.Template.Ref)
	if err != nil {
		return nil, err
	}

	runtime, err := ToDomainRuntime(configDto.Runtime)
	if err != nil {
		return nil, err
	}

	state, err := vos.NewState(configDto.State.Backend, configDto.State.URL)
	if err != nil {
		return nil, err
	}

	auth, err := ToDomainAuth(configDto.Auth)
	if err != nil {
		return nil, err
	}

	return aggregates.NewProject(id, data, template, runtime, state, auth)
}

func ToRuntimeDto(runtime vos.Runtime) dto.RuntimeDTO {
	volumes := make([]dto.VolumeDTO, 0, len(runtime.Volumes()))
	for _, volume := range runtime.Volumes() {
		volumes = append(volumes, dto.VolumeDTO{
			Host: volume.Host(),
			Container: volume.Container(),
		})
	}

	envVars := make([]dto.EnvVarDTO, 0, len(runtime.Env()))
	for _, envVar := range runtime.Env() {
		envVars = append(envVars, dto.EnvVarDTO{
			Name: envVar.Name(),
			Value: envVar.Value(),
		})
	}

	return dto.RuntimeDTO{
		Container: dto.ContainerDTO{
			Image: runtime.Container().Image(),
			Tag: runtime.Container().Tag(),
			CoreVersion: runtime.Container().CoreVersion(),
		},
		Volumes: volumes,
		Env: envVars,
	}
}

func ToAuthDto(auth vos.Auth) dto.AuthDTO {
	extra := make([]map[string]string, 0, len(auth.Params().Extra()))
	for key, value := range auth.Params().Extra() {
		extra = append(extra, map[string]string{key: value})
	}
	return dto.AuthDTO{
		Plugin: auth.Plugin(),
		Params: dto.AuthParamsDTO{
			ClientID: auth.Params().ClientID(),
			GrantType: auth.Params().GrantType(),
			ClientSecret: auth.Params().ClientSecret(),
			Scope: auth.Params().Scope(),
			Extra: extra,
		},
	}
}

func ToDto(config *aggregates.Project) dto.FDConfigDTO {

	projectDto := dto.ProjectDTO{
		ID: config.ID().String(),
		Name: config.Data().Name(),
		Team: config.Data().Team(),
		Description: config.Data().Description(),
		Organization: config.Data().Organization(),
	}

	templateDto := dto.TemplateDTO{
		URL: config.Template().URL(),
		Ref: config.Template().Ref(),
	}

	runtimeDto := ToRuntimeDto(config.Runtime())

	stateDto := dto.StateDTO{
		Backend: config.State().Backend(),
		URL: config.State().URL(),
	}

	authDto := ToAuthDto(config.Auth())

	return dto.FDConfigDTO{
		Project: projectDto,
		Template: templateDto,
		Runtime: runtimeDto,
		State: stateDto,
		Auth: authDto,
	}
}
