package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *AppConfig) handleIndex(resw http.ResponseWriter, req *http.Request) {
	data := generateJson()

	writeJSON(resw, http.StatusAccepted, data)
}

func processAuth(resw http.ResponseWriter, requestPayload *RequestPayload) {
	data, err := authenticate(requestPayload.Auth)
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

func (app *AppConfig) handleProcess(resw http.ResponseWriter, req *http.Request) {
	var requestPayload RequestPayload

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err)
		return
	}

	switch requestPayload.Action {
	case SERVICES.Auth.Name:
		processAuth(resw, &requestPayload)
	default:
		errorJSON(resw, errors.New("unknown Action"))
		return
	}
}
