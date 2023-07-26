package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func generateJson() ResponsePayload {
	payload := ResponsePayload{
		Error:   false,
		Message: "Hit the broker",
	}
	return payload
}

func authenticate(payload AuthPayload) (*ResponsePayload, error) {
	authRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[authenticate] Marshal Error: ", err)
		return nil, err
	}

	authRequestReader := bytes.NewBuffer(authRequest)
	request, err := http.NewRequest(http.MethodPost, SERVICES.Auth.URL, authRequestReader)
	if err != nil {
		log.Println("[authenticate] Request Error: ", err)
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("[authenticate] Client Error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest || response.StatusCode == http.StatusUnauthorized {
		responseErr := errors.New("invalid credentials")
		return nil, responseErr
	}

	if response.StatusCode != http.StatusAccepted {
		responseErr := errors.New("error calling auth service")
		return nil, responseErr
	}

	jsonFromService := ResponsePayload{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&jsonFromService)
	if decodeErr != nil {
		log.Println("[authenticate] Decode Error: ", decodeErr)
		return nil, decodeErr
	}

	return &jsonFromService, nil
}
