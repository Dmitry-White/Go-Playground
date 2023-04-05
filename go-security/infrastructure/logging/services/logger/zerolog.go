package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerZerolog struct {
	logger zerolog.Logger
}

func UseZerolog() ILogger {
	return &LoggerZerolog{
		logger: log.Logger,
	}
}

func (l *LoggerZerolog) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *LoggerZerolog) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *LoggerZerolog) Error(msg string) {
	l.logger.Error().Msg(msg)
}
