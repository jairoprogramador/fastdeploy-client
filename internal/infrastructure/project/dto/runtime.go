package dto

type RuntimeDTO struct {
	Image     string        `yaml:"image"`
	Tag       string        `yaml:"tag"`
	Build     BuildDTO      `yaml:"build,omitempty"`
	Run       RunDTO        `yaml:"run,omitempty"`
}
