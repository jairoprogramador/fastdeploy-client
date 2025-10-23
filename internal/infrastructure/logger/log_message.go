package logger

import (
	"fmt"
	"io"
	"log"

	"github.com/jairoprogramador/fastdeploy/internal/application/ports"
)

type LoggerMessage struct {
	infoLogger   *log.Logger
	detailLogger *log.Logger
}

func NewLoggerService(infoHandle io.Writer, detailHandle io.Writer, showTimestamp bool) ports.LogMessage {
	flags := 0
	if showTimestamp {
		flags = log.LstdFlags
	}
	return &LoggerMessage{
		infoLogger:   log.New(infoHandle, "", flags),
		detailLogger: log.New(detailHandle, "", log.LstdFlags),
	}
}

func (s *LoggerMessage) Info(msg string) {
	s.infoLogger.Println(msg)
}

func (s *LoggerMessage) Detail(msg string) {
	s.detailLogger.Println(msg)
}

func (s *LoggerMessage) Success(msg string) {
	message := fmt.Sprintf("✅ %s", msg)
	s.Info(message)
	s.Detail(message)
}

func (s *LoggerMessage) Error(msg string) {
	message := fmt.Sprintf("❌ %s", msg)
	s.Info(message)
	s.Detail(message)
}
