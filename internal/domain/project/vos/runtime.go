package vos

const (
	DefaultCoreVersion = "1.0.3"
	DefaultImageSource = "fastdeploy/runner-java17-springboot"
	DefaultImageTag    = "latest"
)

const (
	DefaultContainerProjectPath = "/home/fastdeploy/app"
	DefaultContainerStatePath   = "/home/fastdeploy/.fastdeploy"
)

const (
	ProjectPathKey = "projectPath"
	StatePathKey   = "statePath"
)

type Runtime struct {
	CoreVersion string
	Image       Image
	Volumes     []Volume
}

type Image struct {
	Source string
	Tag    string
}

type Volume struct {
	Host      string
	Container string
}
