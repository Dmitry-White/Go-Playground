package main

import (
	"errors"
	"fmt"
	"go-microservices/logger-service/dal"
	"net/http"
)

func (app *AppConfig) handleWriteLog(resw http.ResponseWriter, req *http.Request) {
	requestPayload := RequestPayload{}

	err := readJSON(resw, req, &requestPayload)
	if err != nil {
		errorJSON(resw, err, http.StatusBadRequest)
		return
	}

	logEntry := dal.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	entry, err := app.Models.LogEntry.Insert(logEntry)
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
	entries, err := app.Models.LogEntry.GetAll()
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
