package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *AppConfig) handleMail(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}

	msg, err := app.Services.send(&requestPayload)
	if err != nil {
		errorJSON(resw, errors.New("failed to send email"), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Email is sent to: %s", msg.To),
		Data:    requestPayload,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}
