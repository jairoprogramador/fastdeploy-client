package ports

import "io"

type LogRepository interface {
	GetLatest() ([]byte, error)
	CreateFile() (io.WriteCloser, error)
}
