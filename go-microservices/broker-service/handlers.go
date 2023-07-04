package main

import (
	"fmt"
	"net/http"
)

func (a *Config) handleIndex(resw http.ResponseWriter, req *http.Request) {
	data, jsonErr := generateJson()
	if jsonErr != nil {
		fmt.Fprintln(resw, jsonErr.Error())
	}

	resw.Header().Set("Content-Type", "application/json")
	resw.WriteHeader(http.StatusAccepted)

	_, err := resw.Write(data)
	if err != nil {
		fmt.Fprintln(resw, err.Error())
	}
}
