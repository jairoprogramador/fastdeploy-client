package ports

type CoreVersion interface {
	GetLatestVersion() (string, error)
}