package ports

import (
	"context"
	"github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
)

type DockerService interface {
	Check(ctx context.Context) error
	Build(ctx context.Context, opts vos.BuildOptions) error
	Run(ctx context.Context, opts vos.RunOptions) error
}
