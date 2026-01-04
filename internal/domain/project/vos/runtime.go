package vos

type Runtime struct {
	container Container
	volumes   []Volume
	env       []EnvVar
}

func NewRuntime(container Container, volumes []Volume, env []EnvVar) Runtime {
	return Runtime{container: container, volumes: volumes, env: env}
}

func (r Runtime) Container() Container { return r.container }
func (r Runtime) Volumes() []Volume    { return r.volumes }
func (r Runtime) Env() []EnvVar        { return r.env }
