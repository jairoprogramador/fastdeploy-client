package vos

import "errors"

const (
	DefaultContainerImage = "Dockerfile"
	DefaultContainerTag   = "latest"
)

type Image struct {
	image string
	tag   string
}

func NewImage(image, tag string) (Image, error) {
	if image == "" {
		return Image{}, errors.New("image is required")
	}
	if tag == "" {
		return Image{}, errors.New("tag is required")
	}
	return Image{image: image, tag: tag}, nil
}

func (c Image) Image() string { return c.image }
func (c Image) Tag() string   { return c.tag }
