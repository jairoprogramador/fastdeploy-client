package ports

type UserInputService interface {
	Ask(question, defaultValue string) (string, error)
}
