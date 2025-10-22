package vos

const (
	DefaultStateBackend = "local"
	DefaultStateURL     = ""
)

type State struct {
	Backend string
	URL     string
}
