package logger

import (
	"log"
)

type LoggerCustom struct {
	logger log.Logger
}

func UseCustom() ILogger {
	logger := log.New(log.Writer(), "", log.Flags())
	return &LoggerCustom{
		logger: *logger,
	}
}

func (l *LoggerCustom) Info(msg string) {
	l.logger.Println("Info:", msg)
}

func (l *LoggerCustom) Warn(msg string) {
	l.logger.Println("Warn:", msg)
}

func (l *LoggerCustom) Error(msg string) {
	l.logger.Println("Error:", msg)
}
