package aggregates

import (
	"errors"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
)

type Project struct {
	id        vos.ProjectID
	data      vos.ProjectData
	template  vos.Template
	runtime   vos.Runtime
	state     vos.State
	auth      vos.Auth
}

func NewProject(
	id vos.ProjectID,
	data vos.ProjectData,
	template vos.Template,
	runtime vos.Runtime,
	state vos.State,
	auth vos.Auth) (*Project, error) {
	if template.URL() == "" {
		return nil, errors.New("template is required")
	}
	if template.Ref() == "" {
		return nil, errors.New("template ref is required")
	}
	if runtime.Container().Image() == "" {
		return nil, errors.New("runtime container image is required")
	}
	if runtime.Container().Tag() == "" {
		return nil, errors.New("runtime container tag is required")
	}
	return &Project{
		id: id,
		data: data,
		template: template,
		runtime: runtime,
		state: state,
		auth: auth,
		}, nil
}

func (p *Project) IsIDDirty() bool {
	generatedID := vos.GenerateProjectID(p.data.Name(), p.data.Organization(), p.data.Team())
	if !p.id.Equals(generatedID) {
		p.id = generatedID
		return true
	}
	return false
}

func (p *Project) ID() vos.ProjectID {
	return p.id
}

func (p *Project) Data() vos.ProjectData {
	return p.data
}

func (p *Project) Template() vos.Template {
	return p.template
}

func (p *Project) Runtime() vos.Runtime {
	return p.runtime
}

func (p *Project) State() vos.State {
	return p.state
}

func (p *Project) Auth() vos.Auth {
	return p.auth
}