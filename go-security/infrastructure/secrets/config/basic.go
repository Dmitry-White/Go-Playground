package config

import (
	"fmt"
	"os"
)

const (
	API_KEY = "API_KEY"
	API_URL = "API_URL"
)

type Config struct {
	API_KEY string
	API_URL string
}

func LoadConfig() (*Config, error) {
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
