package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-microservices/authentication-service/dal"
	"log"
	"net/http"
)

func (s *Services) authenticate(requestPayload *RequestPayload) (*dal.User, error) {
	user, err := s.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		log.Println("[authenticate] Model Error: ", err)
		return nil, err
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		log.Println("[authenticate] Password Error: ", err)
		return nil, err
	}

	return user, nil
}

func (s *Services) log(user *dal.User) (*ResponsePayload, error) {
	payload := BrokerRequestPayload{
		Action: "Log",
		Log: LogPayload{
			Name: "Authentication",
			Data: fmt.Sprintf("%s logged in", user.Email),
		},
	}
	logRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[log] Marshal Error: ", err)
		return nil, err
	}

	logRequestReader := bytes.NewBuffer(logRequest)
	request, err := http.NewRequest(http.MethodPost, s.Broker.URL, logRequestReader)
	if err != nil {
		log.Println("[log] Request Error: ", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("[log] Client Error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		responseErr := errors.New("error calling broker service")
		return nil, responseErr
	}

	jsonFromService := ResponsePayload{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&jsonFromService)
	if decodeErr != nil {
		log.Println("[log] Decode Error: ", decodeErr)
		return nil, decodeErr
	}

	return &jsonFromService, nil
}
