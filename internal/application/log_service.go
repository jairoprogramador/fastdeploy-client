package application

import (
	"fmt"
	"io"

	"github.com/jairoprogramador/fastdeploy/internal/application/ports"
)

type LogService struct {
	repo ports.LogRepository
	logMessage ports.LogMessage
}

func NewLogService(repo ports.LogRepository, logMessage ports.LogMessage) *LogService {
	return &LogService{
		repo: repo,
		logMessage: logMessage,
	}
}

func (s *LogService) GetLatestLog() ([]byte, error) {
	content, err := s.repo.GetLatest()
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return nil, err
	}
	if len(content) == 0 {
		s.logMessage.Info("there are no logs yet")
		return []byte{}, nil
	} else {
		s.logMessage.Info(string(content))
	}
	return content, nil
}

func (s *LogService) CreateFile() (io.WriteCloser, error) {
	s.logMessage.Detail("Creating log file...")
	file, err := s.repo.CreateFile()
	if err != nil {
		s.logMessage.Error(fmt.Sprintf("%v", err))
		return nil, err
	}
	s.logMessage.Detail("Log file created successfully")
	return file, nil
}
