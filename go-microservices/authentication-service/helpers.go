package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Reads the body of a request and converts it into JSON
// TODO: Extract MaxBytesReader into a middleware
func readJSON(resw http.ResponseWriter, req *http.Request, data any) error {
	req.Body = http.MaxBytesReader(resw, req.Body, int64(READ_LIMIT))

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// Takes a response status code and arbitrary data and writes a json response to the client
func writeJSON(resw http.ResponseWriter, status int, data any, headers ...http.Header) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			resw.Header()[key] = value
		}
	}

	resw.Header().Set("Content-Type", "application/json")
	resw.WriteHeader(status)

	bytesWritten, err := resw.Write(dataBytes)
	if err != nil {
		return err
	}
	log.Println("[writeJSON] Written:", bytesWritten)

	return nil
}

// Takes an error, and optionally a response status code,
// and generates an JSON error response and sends it as a response
func errorJSON(resw http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := ResponsePayload{
		Error:   true,
		Message: err.Error(),
	}
	log.Println("[errorJSON] Payload:", payload)

	return writeJSON(resw, statusCode, payload)
}
