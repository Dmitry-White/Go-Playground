package main

import (
	"errors"
	"fmt"
	"go-security/infrastructure/logging/services/logger"
	"log"
)

type LoggerFunc func() logger.ILogger

const (
	CUSTOM_STRATEGY  = "custom"
	ZEROLOG_STRATEGY = "zerolog"
	ZAP_STRATEGY     = "zap"
)

func getLoggerFunc(strategy string) (LoggerFunc, error) {
	choice := fmt.Sprintf("Using %s logger", strategy)
	log.Println(choice)

	switch strategy {
	case CUSTOM_STRATEGY:
		return logger.UseCustom, nil
	case ZEROLOG_STRATEGY:
		return logger.UseZerolog, nil
	case ZAP_STRATEGY:
		return logger.UseZap, nil
	}
	return nil, errors.New("this strategy is not supported")
}
