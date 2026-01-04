package dto

type ProjectDTO struct {
	ID           string `yaml:"id"`
	Name         string `yaml:"name"`
	Team         string `yaml:"team"`
	Description  string `yaml:"description,omitempty"`
	Organization string `yaml:"organization"`
}