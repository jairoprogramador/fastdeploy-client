package ports

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	proAgg "github.com/jairoprogramador/fastdeploy/internal/domain/project/aggregates"
)

// ImageService define el contrato para la lógica de construcción de opciones de imagen.
type ImageService interface {
	CreateOptions(project *proAgg.Project) (vos.ImageOptions, error)
	BuildCommand(opts vos.ImageOptions) (string, error)
}