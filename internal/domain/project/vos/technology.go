package vos

const (
	DefaultStack = "springboot"
	DefaultInfrastructure = "azure"
)

type Technology struct {
	Stack          string
	Infrastructure string
}
