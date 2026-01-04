package dto

type FDConfigDTO struct {
	Project  ProjectDTO  `yaml:"project"`
	Template TemplateDTO `yaml:"template"`
	Runtime  RuntimeDTO  `yaml:"runtime"`
	State    StateDTO    `yaml:"state"`
	Auth     AuthDTO     `yaml:"auth,omitempty"`
}





