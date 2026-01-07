package ports

import "context"

type CommandExecutor interface {
	Execute(ctx context.Context, command string) (string, error)
}
