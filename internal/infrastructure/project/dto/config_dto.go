package dto

type FileConfig struct {
	Project struct {
		Name         string `yaml:"name"`
		Version      string `yaml:"version"`
		Team         string `yaml:"team"`
		Description  string `yaml:"description"`
		Organization string `yaml:"organization"`
	} `yaml:"project"`
	Template struct {
		URL string `yaml:"url"`
		Ref string `yaml:"ref"`
	} `yaml:"template"`
	Technology struct {
		Stack          string `yaml:"stack"`
		Infrastructure string `yaml:"infrastructure"`
	} `yaml:"technology"`
	Runtime struct {
		Image struct {
			Source string `yaml:"source"`
			Tag    string `yaml:"tag"`
		} `yaml:"image"`
		Volumes struct {
			ProjectMountPath string `yaml:"project_mount_path"`
			StateMountPath   string `yaml:"state_mount_path"`
		} `yaml:"volumes"`
	} `yaml:"runtime"`
	State struct {
		Backend string `yaml:"backend"`
		URL     string `yaml:"url"`
	} `yaml:"state"`
}
