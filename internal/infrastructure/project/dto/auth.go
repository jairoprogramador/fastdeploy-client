package dto

type AuthDTO struct {
	Plugin string        `yaml:"plugin,omitempty"`
	Params AuthParamsDTO `yaml:"params,omitempty"`
}
