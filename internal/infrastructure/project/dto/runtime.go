package dto

type RuntimeDTO struct {
	Container ContainerDTO `yaml:"container"`
	Volumes   []VolumeDTO  `yaml:"volumes,omitempty"`
	Env       []EnvVarDTO  `yaml:"env,omitempty"`
}
