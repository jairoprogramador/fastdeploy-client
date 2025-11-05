package vos

type Auth struct {
	Plugin string
	Params AuthParams
}

type AuthParams struct {
	ClientID     string
	GrantType    string
	ClientSecret string
	Scope        string
	Extra        map[string]string
}
