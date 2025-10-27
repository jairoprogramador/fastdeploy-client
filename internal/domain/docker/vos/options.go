package vos

type BuildOptions struct {
	Image      Image
	Dockerfile string
	Context    string
	Args       map[string]string
}

type RunOptions struct {
	Image        Image
	Volumes      []Volume
	EnvVars      map[string]string
	Command      string
	Interactive  bool
	AllocateTTY  bool
	RemoveOnExit bool
	WorkDir      string
	Groups       []string
}
