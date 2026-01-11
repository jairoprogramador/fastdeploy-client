package vos

import "errors"

type BuildOptions struct {
	image ImageName
	args  map[string]string
}

func NewBuildOptions(
	image ImageName, args map[string]string) (BuildOptions, error) {

	if image == (ImageName{}) {
		return BuildOptions{}, errors.New("image is required")
	}

	return BuildOptions{
		image: image,
		args:  args,
	}, nil
}

func (b BuildOptions) Image() ImageName {
	return b.image
}

func (b BuildOptions) Args() map[string]string {
	return b.args
}
