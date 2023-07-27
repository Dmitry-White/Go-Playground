package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *AppConfig) handleWriteLog(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}

	entry, err := app.Services.write(&requestPayload)
	if err != nil {
		errorJSON(resw, errors.New("failed to log"), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged entry: %s", entry.InsertedID),
		Data:    entry,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}

func (app *AppConfig) handleReadLogs(resw http.ResponseWriter, req *http.Request) {
	entries, err := app.Services.all()
	if err != nil {
		errorJSON(resw, errors.New("failed to access logs"), http.StatusInternalServerError)
		return
	}

	responsePayload := ResponsePayload{
		Error:   false,
		Message: fmt.Sprintf("Logged entries: %d", len(entries)),
		Data:    entries,
	}
	writeJSON(resw, http.StatusAccepted, responsePayload)
}
