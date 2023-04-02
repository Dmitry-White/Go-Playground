package main

import (
	"context"
	"fmt"
	"go-security/infrastructure/secrets/config"
	"log"
	"net/http"
	"time"
)

var serverConfig *config.Config

func prepareEnv() {
	envFunc, err := getEnvFunc(VIPER_STRATEGY)
	if err != nil {
		log.Fatal(err)
	}

	loadedConfig, err := envFunc()
	if err != nil {
		log.Fatal(err)
	}

	serverConfig = loadedConfig
	log.Printf("Server Config: %+v", serverConfig)
}

func handler(resw http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customRequest, err := http.NewRequestWithContext(ctx, "GET", serverConfig.API_URL, nil)
	if err != nil {
		http.Error(resw, "can't create request", http.StatusInternalServerError)
		return
	}
	customRequest.SetBasicAuth(serverConfig.API_USER, serverConfig.API_KEY)

	resp, err := http.DefaultClient.Do(customRequest)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(resw, "can't call API", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(resw, "OK\n")
}
