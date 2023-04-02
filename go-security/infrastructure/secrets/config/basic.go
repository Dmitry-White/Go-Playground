package config

import (
	"fmt"
	"os"
)

func LoadBasicConfig() (*Config, error) {
	apiKey := os.Getenv(API_KEY)
	if apiKey == "" {
		return nil, fmt.Errorf("missing %s", API_KEY)
	}

	apiUrl := os.Getenv(API_URL)
	if apiUrl == "" {
		return nil, fmt.Errorf("missing %s", API_URL)
	}

	config := &Config{
		API_KEY: apiKey,
		API_URL: apiUrl,
	}

	return config, nil
}
