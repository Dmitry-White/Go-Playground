package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (s *Services) listen(handler HandlerFunc) error {
	messages, err := s.Models.Consumer.Listen()
	if err != nil {
		return err
	}

	log.Println("[listen] Source Exchange:", s.Models.Consumer.Exchange)
	log.Println("[listen] Source Queue:", s.Models.Consumer.Queue.Name)

	forever := make(chan bool)
	go handler(messages)

	log.Println("[listen] Waiting for a message...")
	<-forever

	return nil
}

func (s *Services) log(payload LogPayload) error {
	logRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[log] Marshal Error: ", err)
		return err
	}

	logRequestReader := bytes.NewBuffer(logRequest)
	request, err := http.NewRequest(http.MethodPost, s.Log.URL, logRequestReader)
	if err != nil {
		log.Println("[log] Request Error: ", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("[log] Client Error: ", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		responseErr := errors.New("[log] error calling log service")
		return responseErr
	}

	log.Printf("[log] Logged: %+v\n", payload)

	return nil
}
