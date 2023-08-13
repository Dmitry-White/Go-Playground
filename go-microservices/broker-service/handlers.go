package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *AppConfig) handleSetup() error {
	exchangeErr := app.Services.Models.Emmiter.SetupExchange(EXCHANGES.LOGS)
	if exchangeErr != nil {
		return exchangeErr
	}

	return nil
}

func (app *AppConfig) handleIndex(resw http.ResponseWriter, req *http.Request) {
	data, err := app.Services.index()
	if err != nil {
		errorJSON(resw, err)
		return
	}

	writeJSON(resw, http.StatusAccepted, data)
}

func (app *AppConfig) processAuth(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.authenticate(requestPayload.Auth)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusUnauthorized)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed Auth request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) processLog(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.log(requestPayload.Log)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusUnauthorized)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed Log request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) processMail(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.mail(requestPayload.Mail)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusUnauthorized)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed Mail request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) processAsync(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.async(requestPayload.Async)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusUnauthorized)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed Async request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) processRPC(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.logRPC(requestPayload.LogRPC)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed RPC Log request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) processGRPC(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := app.Services.logGRPC(requestPayload.LogGRPC)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	if data.Error {
		errorJSON(resw, errors.New(data.Message), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Processed GRPC Log request: %s", data.Message),
		Data:    data,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) handleProcess(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err)
		return
	}
	log.Println("[handleProcess] Request: ", requestPayload)

	switch requestPayload.Action {
	case SERVICES.Auth.Name:
		app.processAuth(resw, &requestPayload)
	case SERVICES.Log.Name:
		app.processLog(resw, &requestPayload)
	case SERVICES.Mail.Name:
		app.processMail(resw, &requestPayload)
	case SERVICES.Async.Name:
		app.processAsync(resw, &requestPayload)
	case SERVICES.LogRPC.Name:
		app.processRPC(resw, &requestPayload)
	case SERVICES.LogGRPC.Name:
		app.processGRPC(resw, &requestPayload)
	default:
		errorJSON(resw, errors.New("unknown Action"))
		return
	}
}
