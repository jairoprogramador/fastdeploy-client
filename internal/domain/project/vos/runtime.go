package vos

type Runtime struct {
	image Image
	volumes   []Volume
	env       []EnvVar
	args      []Argument
}

func NewRuntime(image Image, volumes []Volume, env []EnvVar, args []Argument) Runtime {
	return Runtime{image: image, volumes: volumes, env: env, args: args}
}

func (r Runtime) Image() Image  { return r.image }
func (r Runtime) Volumes() []Volume { return r.volumes }
func (r Runtime) Env() []EnvVar     { return r.env }
func (r Runtime) Args() []Argument  { return r.args }
