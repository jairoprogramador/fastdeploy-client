package dto

type AuthParamsDTO struct {
	ClientID     string              `yaml:"client_id"`
	ClientSecret string              `yaml:"client_secret"`
	GrantType    string              `yaml:"grant_type"`
	Scope        string              `yaml:"scope"`
	Extra        []map[string]string `yaml:"extra,omitempty"`
}