// Package project provides infrastructure implementations for the project domain.
package project

import (
	"os"
	"path/filepath"

	"github.com/jairoprogramador/fastdeploy/internal/domain/project/ports"
	"github.com/jairoprogramador/fastdeploy/internal/domain/project/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/dto"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/project/mapper"
	"gopkg.in/yaml.v3"
)

const configFileName = "fdconfig.yaml"

type yamlProjectRepository struct {
	workDir string
}

func NewYAMLProjectRepository(workDir string) ports.ProjectRepository {
	return &yamlProjectRepository{workDir: workDir}
}

func (r *yamlProjectRepository) Save(config *vos.Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath(), data, 0644)
}

func (r *yamlProjectRepository) Exists() (bool, error) {
	_, err := os.Stat(r.filePath())
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (r *yamlProjectRepository) Load() (*vos.Config, error) {
	data, err := os.ReadFile(r.filePath())
	if err != nil {
		if os.IsNotExist(err) {
			return mapper.ToDomain(dto.FileConfig{}), nil
		}
		return nil, err
	}

	var fileConfig dto.FileConfig
	if err := yaml.Unmarshal(data, &fileConfig); err != nil {
		return nil, err
	}

	return mapper.ToDomain(fileConfig), nil
}

func (r *yamlProjectRepository) filePath() string {
	return filepath.Join(r.workDir, configFileName)
}
