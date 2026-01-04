package ports

import (
	"context"

	"github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
)

type DockerService interface {
	Check(ctx context.Context) error
	Build(ctx context.Context, opts vos.ImageOptions) error
	Run(ctx context.Context, opts vos.ContainerOptions) (string, error)
}
