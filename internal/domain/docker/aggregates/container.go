package aggregates

import (
	"fmt"
	"runtime"
	"strings"

	proAgg "github.com/jairoprogramador/fastdeploy/internal/domain/project/aggregates"
	proVos "github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"

	"github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
)

// Execution es la Raíz del Agregado para una ejecución de Docker.
// Encapsula la lógica para preparar y definir una única ejecución de un comando.
type Container struct {
	buildOpts vos.ImageOptions
	runOpts   vos.ContainerOptions
}

// NewExecution es la fábrica para crear un nuevo agregado Execution.
// Aquí reside la lógica de negocio para construir las opciones de Docker.
func NewContainer(
	project *proAgg.Project,
	command string,
	environment string,
	hostProjectPath string,
	fastdeployHome string,
	isTerminal bool,
	withTty bool,
	coreVersion string,
) (*Container, error) {

	var localImage vos.ImageName
	var buildOpts vos.ImageOptions

	projectRuntime := project.Runtime()
	projectTemplate := project.Template()
	projectData := project.Data()
	containerInfo := projectRuntime.Container()

	// 1. Lógica de "Build vs. Pull"
	if containerInfo.Image() != vos.DefaultDockerfile {
		img, err := vos.NewImageName(containerInfo.Image(), containerInfo.Tag())
		if err != nil {
			return nil, err
		}
		localImage = img
	} else {
		// Preparamos las opciones de build
		localImageName := fmt.Sprintf("%s-%s", projectData.Team(), projectTemplate.DirName())
		img, err := vos.NewImageName(localImageName, containerInfo.Tag())
		if err != nil {
			return nil, err
		}
		localImage = img

		buildArgs := make(map[string]string)
		if runtime.GOOS == "linux" {
			buildArgs["DEV_GID"] = "$(id -g)"
		}

		if containerInfo.CoreVersion() != "" {
			buildArgs["FASTDEPLOY_VERSION"] = containerInfo.CoreVersion()
		} else if coreVersion != "" {
			buildArgs["FASTDEPLOY_VERSION"] = coreVersion
		} else {
			buildArgs["FASTDEPLOY_VERSION"] = vos.DefaultContainerCoreVersion
		}

		buildOpts = vos.ImageOptions{
			image:      localImage,
			context:    ".",
			dockerfile: vos.DefaultDockerfile,
			args:       buildArgs,
		}
	}

	// 2. Preparamos las opciones de Run
	volumesMap := make(map[string]string)
	for _, volume := range projectRuntime.Volumes() {
		volumesMap[volume.Host()] = volume.Container()
	}
	volumesMap[hostProjectPath] = vos.DefaultContainerProjectPath

	envVars := make(map[string]string)
	for _, envVar := range projectRuntime.Env() {
		envVars[envVar.Name()] = envVar.Value() // La resolución de variables se hará en la capa de aplicación
	}

	if project.State().Backend() == proVos.DefaultStateBackend {
		volumesMap[fastdeployHome] = vos.DefaultContainerFastdeployPath
		envVars["FASTDEPLOY_HOME"] = volumesMap[fastdeployHome]
	}

	volumes := make([]vos.Volume, 0, len(volumesMap))
	for hostPath, containerPath := range volumesMap {
		vol, err := vos.NewVolume(hostPath, containerPath)
		if err != nil {
			return nil, err
		}
		volumes = append(volumes, vol)
	}

	allocateTty := withTty && isTerminal
	interactive := allocateTty

	groups := []string{}
	if runtime.GOOS == "linux" {
		groups = append(groups, "$(getent group docker | cut -d: -f3)")
	}

	runOpts := vos.ContainerOptions{
		Image:        localImage,
		Volumes:      volumes,
		EnvVars:      envVars,
		Command:      strings.TrimSpace(fmt.Sprintf("%s %s", command, environment)),
		Interactive:  interactive,
		AllocateTTY:  allocateTty,
		RemoveOnExit: true,
		Groups:       groups,
	}

	return &Execution{
		buildOpts: buildOpts,
		runOpts:   runOpts,
	}, nil
}


// BuildOptions devuelve las opciones de build. Puede ser nil.
func (e *Container) BuildOptions() vos.ImageOptions {
	return e.buildOpts
}

// RunOptions devuelve las opciones de run.
func (e *Container) RunOptions() vos.ContainerOptions {
	return e.runOpts
}
