package main

import (
	"go-security/infrastructure/secrets/config"
)

func prepareBasic() (*config.Config, error) {
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return loadedConfig, nil
}
