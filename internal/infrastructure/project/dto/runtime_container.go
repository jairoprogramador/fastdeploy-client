package dto

type ContainerDTO struct {
	Image       string `yaml:"image,omitempty"`
	Tag         string `yaml:"tag,omitempty"`
	CoreVersion string `yaml:"core_version,omitempty"`
}
