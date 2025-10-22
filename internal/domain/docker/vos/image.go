package vos

import "fmt"

type Image struct {
	Name string
	Tag  string
}

func (image Image) FullName() string {
	return fmt.Sprintf("%s:%s", image.Name, image.Tag)
}