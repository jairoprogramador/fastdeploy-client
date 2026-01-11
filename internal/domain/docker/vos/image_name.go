package vos

import (
	"errors"
	"fmt"
)

type ImageName struct {
	name string
	tag  string
}

func NewImageName(name, tag string) (ImageName, error) {
	if name == "" {
		return ImageName{}, errors.New("image name cannot be empty")
	}
	if tag == "" {
		return ImageName{}, errors.New("image tag cannot be empty")
	}
	return ImageName{name: name, tag: tag}, nil
}

func (i ImageName) Name() string {
	return i.name
}

func (i ImageName) Tag() string {
	return i.tag
}

func (i ImageName) FullName() string {
	return fmt.Sprintf("%s:%s", i.name, i.tag)
}
