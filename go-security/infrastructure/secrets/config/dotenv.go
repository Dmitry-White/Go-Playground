package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotenvConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	apiUser := os.Getenv(API_USER)
	if apiUser == "" {
		return nil, fmt.Errorf("missing %s", API_USER)
	}

	apiKey := os.Getenv(API_KEY)
	if apiKey == "" {
		return nil, fmt.Errorf("missing %s", API_KEY)
	}

	apiUrl := os.Getenv(API_URL)
	if apiUrl == "" {
		return nil, fmt.Errorf("missing %s", API_URL)
	}

	config := &Config{
		API_USER: apiUser,
		API_KEY:  apiKey,
		API_URL:  apiUrl,
	}

	return config, nil
}
