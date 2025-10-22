package vos

const (
	DefaultImageSource = "fastdeploy/runner-java17-springboot"
	DefaultImageTag    = "latest"
)

const (
	DefaultProjectMountPath = "/home/fastdeploy/app"
	DefaultStateMountPath   = "/home/fastdeploy/.fastdeploy"
)

type Runtime struct {
	Image   Image
	Volumes Volumes
}

type Image struct {
	Source string
	Tag    string
}

type Volumes struct {
	ProjectMountPath string
	StateMountPath   string
}
