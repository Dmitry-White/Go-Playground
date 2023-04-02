package main

import (
	"go-security/infrastructure/secrets/config"
	"log"
	"net/http"
)

var serverConfig *config.Config

func prepareEnv() {
	envFunc, err := getEnvFunc(DOTENV_STRATEGY)
	if err != nil {
		log.Fatal(err)
	}

	loadedConfig, err := envFunc()
	if err != nil {
		log.Fatal(err)
	}

	serverConfig = loadedConfig
}

func handler(resw http.ResponseWriter, req *http.Request) {
	log.Println("Not Implemented")
	log.Println("Config: ", serverConfig)
}
