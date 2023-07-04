package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *Config) handleIndex(resw http.ResponseWriter, req *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	out, jsonErr := json.MarshalIndent(payload, "", "\t")
	if jsonErr != nil {
		fmt.Fprintln(resw, jsonErr.Error())
	}

	resw.Header().Set("Content-Type", "application/json")
	resw.WriteHeader(http.StatusAccepted)
	_, err := resw.Write(out)
	if err != nil {
		fmt.Fprintln(resw, err.Error())
	}
}
