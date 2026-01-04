package vos

import "errors"

const (
	DefaultContainerImage  = "Dockerfile"
	DefaultContainerTag    = "latest"
	DefaultContainerCoreVersion = "latest"
)

type Container struct {
	image       string
	tag         string
	coreVersion string
}

func NewContainer(image, tag, coreVersion string) (Container, error) {
	if image == "" {
		return Container{}, errors.New("image is required")
	}
	if tag == "" {
		return Container{}, errors.New("tag is required")
	}
	if coreVersion == "" {
		return Container{}, errors.New("coreVersion is required")
	}
	return Container{image: image, tag: tag, coreVersion: coreVersion}, nil
}

func (c Container) Image() string { return c.image }
func (c Container) Tag() string { return c.tag }
func (c Container) CoreVersion() string { return c.coreVersion }