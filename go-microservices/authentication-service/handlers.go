package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *AppConfig) handleAuth(resw http.ResponseWriter, req *http.Request) {
	var requestPayload RequestPayload

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		errorJSON(resw, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		errorJSON(resw, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	responsePayload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user: %s", user.Email),
		Data:    user,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}
