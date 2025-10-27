package dto

type FileConfig struct {
	Project  ProjectDTO  `yaml:"project"`
	Template TemplateDTO `yaml:"template"`
	Runtime  RuntimeDTO  `yaml:"runtime"`
	State    StateDTO    `yaml:"state"`
	Auth     AuthDTO     `yaml:"auth"`
}

type ProjectDTO struct {
	ID           string `yaml:"id"`
	Name         string `yaml:"name"`
	Version      string `yaml:"version"`
	Team         string `yaml:"team"`
	Description  string `yaml:"description"`
	Organization string `yaml:"organization"`
}

type TemplateDTO struct {
	URL string `yaml:"url"`
	Ref string `yaml:"ref"`
}

type RuntimeDTO struct {
	CoreVersion string      `yaml:"core_version"`
	Image       ImageDTO    `yaml:"image"`
	Volumes     []VolumeDTO `yaml:"volumes,omitempty"`
	Env         []EnvVarDTO `yaml:"env,omitempty"`
}

type StateDTO struct {
	Backend string `yaml:"backend"`
	URL     string `yaml:"url"`
}

type AuthDTO struct {
	Plugin string        `yaml:"plugin"`
	Params AuthParamsDTO `yaml:"params"`
}

type AuthParamsDTO struct {
	ClientID  string              `yaml:"client_id"`
	ClientSecret string           `yaml:"client_secret"`
	GrantType string              `yaml:"grant_type"`
	Scope     string              `yaml:"scope"`
	Extra     []map[string]string `yaml:"extra,omitempty"`
}

type EnvVarDTO struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type ImageDTO struct {
	Source string `yaml:"source,omitempty"`
	Tag    string `yaml:"tag"`
}

type VolumeDTO struct {
	Host      string `yaml:"host"`
	Container string `yaml:"container"`
}
