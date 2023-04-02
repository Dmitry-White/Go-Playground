package main

import (
	"errors"
	"fmt"
	"go-security/infrastructure/secrets/config"
	"log"
)

type EnvFunc func() (*config.Config, error)

const (
	BASIC_STRATEGY  = "basic"
	DOTENV_STRATEGY = "dotenv"
	VIPER_STRATEGY  = "viper"
)

func getEnvFunc(strategy string) (EnvFunc, error) {
	choice := fmt.Sprintf("Using %s env", strategy)
	log.Println(choice)

	switch strategy {
	case BASIC_STRATEGY:
		return prepareBasic, nil
	case DOTENV_STRATEGY:
		return prepareDotenv, nil
	case VIPER_STRATEGY:
		return prepareViper, nil
	}
	return nil, errors.New("this strategy is not supported")
}
