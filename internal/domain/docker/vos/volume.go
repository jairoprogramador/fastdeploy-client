package vos

import "errors"

type Volume struct {
	hostPath      string
	containerPath string
}

func NewVolume(hostPath, containerPath string) (Volume, error) {
	if hostPath == "" {
		return Volume{}, errors.New("volume host path cannot be empty")
	}
	if containerPath == "" {
		return Volume{}, errors.New("volume container path cannot be empty")
	}
	return Volume{hostPath: hostPath, containerPath: containerPath}, nil
}

func (v Volume) HostPath() string {
	return v.hostPath
}

func (v Volume) ContainerPath() string {
	return v.containerPath
}
