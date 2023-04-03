package services

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

func UseCustom() {
	fmt.Println("Custom")
}

func UseZerolog() {
	log.Info().Msg("Zerolog")
}

func UseZap() {
	logger, _ := zap.NewProduction()
	logger.Info("Zap")
}
