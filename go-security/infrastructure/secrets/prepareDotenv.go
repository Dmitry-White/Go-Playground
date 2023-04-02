package main

import (
	"go-security/infrastructure/secrets/config"
)

func prepareDotenv() (*config.Config, error) {
	loadedConfig, err := config.LoadDotenvConfig()
	if err != nil {
		return nil, err
	}

	return loadedConfig, nil
}
