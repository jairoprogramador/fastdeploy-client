package vos

import "errors"

type ContainerOptions struct {
	image        ImageName
	volumes      []Volume
	envVars      map[string]string
	command      string
	interactive  bool
	allocateTTY  bool
	removeOnExit bool
	workDir      string
	groups       []string
}

func NewContainerOptions(
	image ImageName,
	volumes []Volume,
	envVars map[string]string,
	command string,
	interactive bool,
	allocateTTY bool,
	removeOnExit bool,
	workDir string,
	groups []string) (ContainerOptions, error) {

	if image == (ImageName{}) {
		return ContainerOptions{}, errors.New("image is required")
	}
	if command == "" {
		return ContainerOptions{}, errors.New("command is required")
	}

	return ContainerOptions{
		image:        image,
		volumes:      volumes,
		envVars:      envVars,
		command:      command,
		interactive:  interactive,
		allocateTTY:  allocateTTY,
		removeOnExit: removeOnExit,
		workDir:      workDir,
		groups:       groups,
	}, nil
}

func (r ContainerOptions) Image() ImageName {
	return r.image
}

func (r ContainerOptions) Volumes() []Volume {
	return r.volumes
}

func (r ContainerOptions) EnvVars() map[string]string {
	return r.envVars
}

func (r ContainerOptions) Command() string {
	return r.command
}

func (r ContainerOptions) Interactive() bool {
	return r.interactive
}

func (r ContainerOptions) AllocateTTY() bool {
	return r.allocateTTY
}

func (r ContainerOptions) RemoveOnExit() bool {
	return r.removeOnExit
}

func (r ContainerOptions) WorkDir() string {
	return r.workDir
}

func (r ContainerOptions) Groups() []string {
	return r.groups
}
