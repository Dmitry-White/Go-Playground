package logger

import (
	"go.uber.org/zap"
)

type LoggerZap struct {
	logger *zap.Logger
}

func UseZap() ILogger {
	logger, _ := zap.NewProduction()
	return &LoggerZap{
		logger,
	}
}

func (l *LoggerZap) Info(msg string) {
	l.logger.Info(msg)
}

func (l *LoggerZap) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *LoggerZap) Error(msg string) {
	l.logger.Error(msg)
}
