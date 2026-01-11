package ports

import "github.com/jairoprogramador/fastdeploy-client/internal/domain/project/aggregates"

type ProjectRepository interface {
	Save(project *aggregates.Project) error
	Exists() (bool, error)
	Load() (*aggregates.Project, error)
}
