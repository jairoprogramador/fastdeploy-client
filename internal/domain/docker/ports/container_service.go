package ports

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	proAgg "github.com/jairoprogramador/fastdeploy/internal/domain/project/aggregates"
)

// ContainerService define el contrato para la lógica de construcción de opciones de contenedor.
type ContainerService interface {
	CreateOptions(project *proAgg.Project, commandFastdeploy string, image vos.ImageName) (vos.RunOptions, error)
	BuildCommand(opts vos.RunOptions) (string, error)
}
