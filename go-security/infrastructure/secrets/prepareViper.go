package main

import (
	"go-security/infrastructure/secrets/config"
)

func prepareViper() (*config.Config, error) {
	loadedConfig, err := config.LoadViperConfig()
	if err != nil {
		return nil, err
	}

	return loadedConfig, nil
}
