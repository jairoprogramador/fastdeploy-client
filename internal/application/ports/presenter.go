package ports

import (
	"github.com/jairoprogramador/fastdeploy-client/internal/domain/logger/aggregates"
)

type Presenter interface {
	Render(log *aggregates.Logger)
}
