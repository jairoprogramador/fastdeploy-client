package services

import (
	"fmt"
	"runtime"
	"strings"

	docPor "github.com/jairoprogramador/fastdeploy/internal/domain/docker/ports"
	docVos "github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	proAgg "github.com/jairoprogramador/fastdeploy/internal/domain/project/aggregates"
)


// imageBuilder es la implementación del servicio de dominio.
type imageBuilder struct{}

// NewImageBuilder crea una nueva instancia del servicio.
func NewImageBuilder() docPor.ImageService {
	return &imageBuilder{}
}

// CreateImageOptions encapsula la lógica de negocio para determinar cómo se debe construir una imagen.
func (s *imageBuilder) CreateOptions(project *proAgg.Project) (docVos.ImageOptions, error) {
	localImageName := fmt.Sprintf("%s%s", project.Data().Name(), project.ID().String()[0:6])
	imgName, err := docVos.NewImageName(localImageName, project.Runtime().Container().Tag())
	if err != nil {
		return docVos.ImageOptions{}, err
	}

	imgArgs := make(map[string]string)

	if runtime.GOOS == "linux" {
		imgArgs["DEV_GID"] = "$(id -g)"
	}
	imgArgs["FASTDEPLOY_VERSION"] = project.Runtime().Container().CoreVersion()

	return docVos.NewImageOptions(imgName, project.Runtime().Container().Image(), imgArgs)
}

func (s *imageBuilder) BuildCommand(opts docVos.ImageOptions) (string, error) {
	var commandBuilder strings.Builder
	commandBuilder.WriteString("docker build")

	for key, val := range opts.Args() {
		commandBuilder.WriteString(fmt.Sprintf(" --build-arg %s=%s", key, val))
	}

	commandBuilder.WriteString(fmt.Sprintf(" -t %s", opts.Image().FullName()))

	commandBuilder.WriteString(fmt.Sprintf(" -f %s", opts.FileName()))

	commandBuilder.WriteString(fmt.Sprintf(" %s", "."))

	return commandBuilder.String(), nil
}