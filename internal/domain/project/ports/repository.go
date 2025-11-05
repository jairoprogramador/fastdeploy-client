package ports

import "github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"

type ProjectRepository interface {
	Save(config *vos.Config) error
	Exists() (bool, error)
	Load() (*vos.Config, error)
}
