package vos

import (
	"crypto/sha256"
	"fmt"
	"errors"
)

type ProjectID struct {
	value string
}

func NewProjectID(id string) (ProjectID, error) {
	if id == "" {
		return ProjectID{}, errors.New("id is required")
	}
	return ProjectID{value: id}, nil
}

func GenerateProjectID(name, organization, team string) ProjectID {
	data := fmt.Sprintf("%s-%s-%s", name, organization, team)
	hash := sha256.Sum256([]byte(data))
	return ProjectID{value: fmt.Sprintf("%x", hash)}
}

func (p ProjectID) String() string {
	return p.value
}

func (p ProjectID) Equals(other ProjectID) bool {
	return p.value == other.value
}