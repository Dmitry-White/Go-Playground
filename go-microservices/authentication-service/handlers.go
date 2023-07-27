package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *AppConfig) handleAuth(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}
	log.Println("[handleAuth] Request:", requestPayload)

	user, err := app.Services.authenticate(&requestPayload)
	if err != nil {
		errorJSON(resw, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	logResult, err := app.Services.log(user)
	if err != nil {
		log.Println("[handleAuth] Log Error:", err)
		errorJSON(resw, errors.New("logging issue"), http.StatusInternalServerError)
		return
	}
	log.Println("[handleAuth] Payload:", logResult)

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged in user: %s", user.Email),
		Data:    user,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}
