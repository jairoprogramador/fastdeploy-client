package ports

type VarsResolver interface {
	Resolve(input string, internalVars map[string]string) string
}