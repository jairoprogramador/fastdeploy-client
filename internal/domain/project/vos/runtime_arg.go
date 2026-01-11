package vos

import "errors"

type Argument struct {
	name  string
	value string
}

func NewArgument(name, value string) (Argument, error) {
	if name == "" {
		return Argument{}, errors.New("name is required")
	}
	if value == "" {
		return Argument{}, errors.New("value is required")
	}
	return Argument{name: name, value: value}, nil
}

func (a Argument) Name() string  { return a.name }
func (a Argument) Value() string { return a.value }
