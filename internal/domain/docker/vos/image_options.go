package vos

import "errors"

type ImageOptions struct {
	image        ImageName
	fileName     string
	args         map[string]string
}

func NewImageOptions(
	image ImageName, fileName string, args map[string]string) (ImageOptions, error) {

	if image == (ImageName{}) {
		return ImageOptions{}, errors.New("image is required")
	}

	return ImageOptions{
		image:      image,
		fileName:   fileName,
		args:       args,
	}, nil
}

func (b ImageOptions) Image() ImageName {
	return b.image
}

func (b ImageOptions) Args() map[string]string {
	return b.args
}

func (b ImageOptions) FileName() string {
	return b.fileName
}