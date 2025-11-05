package vos

const (
	DefaultCoreVersion = "1.0.6"
	DefaultImageSource = "jairoprogramador/fdrunjava17azure"
	DefaultImageTag    = "latest"
)

const (
	defaultContainerHomePath       = "/home/fastdeploy"
	DefaultContainerProjectPath    = defaultContainerHomePath + "/app"
	DefaultContainerFastdeployPath = defaultContainerHomePath + "/.fastdeploy"
)

const (
	ProjectPathKey = "projectPath"
	StatePathKey   = "statePath"
)

type Runtime struct {
	CoreVersion string
	Image       Image
	Volumes     []Volume
	Env         []EnvVar
}

type Image struct {
	Source string
	Tag    string
}

type Volume struct {
	Host      string
	Container string
}

type EnvVar struct {
	Name  string
	Value string
}
