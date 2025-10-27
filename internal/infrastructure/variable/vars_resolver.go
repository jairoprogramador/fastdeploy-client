package variable

import (
	"os"
	"strings"

	appPor "github.com/jairoprogramador/fastdeploy/internal/application/ports"
)

type VariableResolver struct{}

func NewVariableResolver() appPor.VarsResolver {
	return &VariableResolver{}
}

func (s *VariableResolver) Resolve(input string, internalVars map[string]string) string {
	if !strings.HasPrefix(input, "{") || !strings.HasSuffix(input, "}") {
		return input
	}

	trimmed := strings.Trim(input, "{}")
	parts := strings.SplitN(trimmed, ".", 2)
	if len(parts) != 2 {
		return input
	}

	sourceType := parts[0]
	key := parts[1]

	switch sourceType {
	case "env":
		return os.Getenv(key)
	case "var":
		if val, ok := internalVars[key]; ok {
			return val
		}
		return ""
	default:
		return input
	}
}
