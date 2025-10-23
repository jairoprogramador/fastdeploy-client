package logger

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jairoprogramador/fastdeploy/internal/application/ports"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/path"
)

type FileLogRepository struct {
	pathResolver *path.PathService
}

func NewFileLogRepository(pathResolver *path.PathService) ports.LogRepository {
	return &FileLogRepository{
		pathResolver: pathResolver,
	}
}

func (r *FileLogRepository) GetLatest() ([]byte, error) {
	logDir := r.pathResolver.GetLogsPath()
	files, err := os.ReadDir(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil
		}
		return nil, err
	}

	var logFiles []fs.DirEntry
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "exec-") && strings.HasSuffix(file.Name(), ".log") {
			logFiles = append(logFiles, file)
		}
	}

	if len(logFiles) == 0 {
		return []byte{}, nil
	}

	sort.Slice(logFiles, func(i, j int) bool {
		return logFiles[i].Name() > logFiles[j].Name()
	})

	latestLogFile := filepath.Join(logDir, logFiles[0].Name())
	content, err := os.ReadFile(latestLogFile)
	if err != nil {
		return nil, fmt.Errorf("error reading log file '%s': %w", latestLogFile, err)
	}

	return content, nil
}

func (r *FileLogRepository) CreateFile() (io.WriteCloser, error) {
	err := r.DeleteFiles()
	if err != nil {
		return nil, err
	}
	logDir := r.pathResolver.GetLogsPath()
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}
	logFileName := fmt.Sprintf("exec-%s.log", time.Now().Format("20060102150405"))
	logFilePath := filepath.Join(logDir, logFileName)
	return os.Create(logFilePath)
}

func (r *FileLogRepository) DeleteFiles() error {
	logDir := r.pathResolver.GetLogsPath()
	files, err := os.ReadDir(logDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "exec-") && strings.HasSuffix(file.Name(), ".log") {
			err = os.Remove(filepath.Join(logDir, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}