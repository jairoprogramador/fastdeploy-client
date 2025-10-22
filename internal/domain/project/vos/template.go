package vos

const (
	DefaultUrl = "https://github.com/jairoprogramador/mydeploytest.git"
	DefaultRef = "main"
)

type Template struct {
	URL string
	Ref string
}
