package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-microservices/broker-service/grpc"
	"log"
	"net/http"
	"net/rpc"
	"time"

	grpcLib "google.golang.org/grpc"
	grpcLibInsecure "google.golang.org/grpc/credentials/insecure"
)

func (s *Services) index() (*ResponsePayload, error) {
	payload := ResponsePayload{
		Error:   false,
		Message: "Hit the broker",
	}

	return &payload, nil
}

func (s *Services) authenticate(payload AuthPayload) (*ResponsePayload, error) {
	authRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[authenticate] Marshal Error: ", err)
		return nil, err
	}

	authRequestReader := bytes.NewBuffer(authRequest)
	request, err := http.NewRequest(http.MethodPost, s.Auth.URL, authRequestReader)
	if err != nil {
		log.Println("[authenticate] Request Error: ", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

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

func (s *Services) log(payload LogPayload) (*ResponsePayload, error) {
	logRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[log] Marshal Error: ", err)
		return nil, err
	}

	logRequestReader := bytes.NewBuffer(logRequest)
	request, err := http.NewRequest(http.MethodPost, s.Log.URL, logRequestReader)
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
		responseErr := errors.New("[log] error calling log service")
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

func (s *Services) mail(payload MailPayload) (*ResponsePayload, error) {
	mailRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[mail] Marshal Error: ", err)
		return nil, err
	}

	mailRequestReader := bytes.NewBuffer(mailRequest)
	request, err := http.NewRequest(http.MethodPost, s.Mail.URL, mailRequestReader)
	if err != nil {
		log.Println("[mail] Request Error: ", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("[mail] Client Error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		responseErr := errors.New("error calling mail service")
		return nil, responseErr
	}

	jsonFromService := ResponsePayload{}
	decoder := json.NewDecoder(response.Body)
	decodeErr := decoder.Decode(&jsonFromService)
	if decodeErr != nil {
		log.Println("[mail] Decode Error: ", decodeErr)
		return nil, decodeErr
	}

	return &jsonFromService, nil
}

func (s *Services) async(payload AsyncPayload) (*ResponsePayload, error) {
	asyncRequest, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("[async] Marshal Error: ", err)
		return nil, err
	}

	pushErr := s.Models.Emmiter.Push(TOPICS.LOG_INFO, asyncRequest)
	if pushErr != nil {
		log.Println("[async] Push Error: ", pushErr)
		return nil, pushErr
	}

	jsonFromService := ResponsePayload{
		Error:   false,
		Message: "Processing started",
	}

	return &jsonFromService, nil
}

func (s *Services) logRPC(payload LogRPCPayload) (*ResponsePayload, error) {
	log.Printf("[logRPC] RPC Log Config: %+v\n", s.LogRPC)

	client, err := rpc.Dial("tcp", s.LogRPC.URL)
	if err != nil {
		log.Println("[logRPC] Client Error: ", err)
		return nil, err
	}

	jsonFromService := ResponsePayload{}
	requestErr := client.Call("AppConfig.HandleLogRPC", &payload, &jsonFromService)
	if requestErr != nil {
		log.Println("[logRPC] Request Error: ", requestErr)
		return nil, requestErr
	}

	return &jsonFromService, nil
}

func (s *Services) logGRPC(payload LogGRPCPayload) (*ResponsePayload, error) {
	log.Printf("[logGRPC] GRPC Log Config: %+v\n", s.LogGRPC)

	blockingReqOption := grpcLib.WithBlock()
	credentials := grpcLibInsecure.NewCredentials()
	credentialsOption := grpcLib.WithTransportCredentials(credentials)

	connection, err := grpcLib.Dial(s.LogGRPC.URL, credentialsOption, blockingReqOption)
	if err != nil {
		log.Println("[logGRPC] Connection Error: ", err)
		return nil, err
	}
	defer connection.Close()

	client := grpc.NewLogServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	requestPayload := &grpc.LogRequest{
		Input: &grpc.Log{
			Name: payload.Name,
			Data: payload.Data,
		},
	}
	response, err := client.HandleLogGRPC(ctx, requestPayload)
	if err != nil {
		log.Println("[logGRPC] Handler Error: ", err)
		return nil, err
	}

	jsonFromService := ResponsePayload{
		Error:   response.Error,
		Message: response.Message,
		Data:    response.Data,
	}

	return &jsonFromService, nil
}
