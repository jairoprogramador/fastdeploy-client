package project

import (
	"crypto/sha256"
	"fmt"
	"sort"

	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/services"
)

type ShaGeneratorID struct{}

func NewShaGeneratorID() services.GeneratorID {
	return &ShaGeneratorID{}
}

func (g *ShaGeneratorID) ProjectID(config *vos.Config) string {
	fields := []string{
		fmt.Sprintf("template:%s", config.Template.URL),
		fmt.Sprintf("stack:%s", config.Technology.Stack),
		fmt.Sprintf("infrastructure:%s", config.Technology.Infrastructure),
	}
	sort.Strings(fields)

	var combined string
	for _, field := range fields {
		combined += field + "|"
	}

	hash := sha256.Sum256([]byte(combined))
	return fmt.Sprintf("%x", hash)
}
