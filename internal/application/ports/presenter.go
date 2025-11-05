package ports

import (
	"github.com/jairoprogramador/fastdeploy/internal/domain/logger/aggregates"
)

type Presenter interface {
	Render(log *aggregates.Logger)
}
