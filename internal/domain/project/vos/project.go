package vos

const (
	DefaultProjectVersion = "1.0.0"
	DefaultProjectTeam    = "shikigami"
	DefaultProjectDescription = "mi despliegue con fastdeploy"
	DefaultProjectOrganization = "fastdeploy"
)

type Project struct {
	Name         string
	Version      string
	Team         string
	Description  string
	Organization string
}
