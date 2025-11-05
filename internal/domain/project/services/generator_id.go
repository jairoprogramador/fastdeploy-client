package services

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
)

type GeneratorID interface {
	ProjectID(config *vos.Config) string
}
