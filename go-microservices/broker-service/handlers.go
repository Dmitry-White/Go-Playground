package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *AppConfig) handleIndex(resw http.ResponseWriter, req *http.Request) {
	data := generateJson()

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
	default:
		errorJSON(resw, errors.New("unknown Action"))
		return
	}
}
