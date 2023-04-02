package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadViperConfig() (*Config, error) {
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	/*
		viper.Get() returns an empty interface{}
		to get the underlying type of the key,
		we have to do the type assertion, we know the underlying value is string
		if we type assert to other type it will throw an error
		If the type is a string then ok will be true
		ok will make sure the program not break
	*/
	apiUser, ok := viper.Get(API_USER).(string)
	if !ok {
		return nil, fmt.Errorf("missing %s", API_USER)
	}

	apiKey, ok := viper.Get(API_KEY).(string)
	if !ok {
		return nil, fmt.Errorf("missing %s", API_KEY)
	}

	apiUrl, ok := viper.Get(API_URL).(string)
	if !ok {
		return nil, fmt.Errorf("missing %s", API_URL)
	}

	config := &Config{
		API_USER: apiUser,
		API_KEY:  apiKey,
		API_URL:  apiUrl,
	}

	return config, nil
}
