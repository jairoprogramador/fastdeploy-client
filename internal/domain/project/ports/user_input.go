package ports

import "context"

type UserInputService interface {
	Ask(ctx context.Context, question, defaultValue string) (string, error)
}
