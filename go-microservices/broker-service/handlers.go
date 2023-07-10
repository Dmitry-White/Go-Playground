package main

import (
	"net/http"
)

func (a *AppConfig) handleIndex(resw http.ResponseWriter, req *http.Request) {
	data := generateJson()

	writeJSON(resw, http.StatusAccepted, data)
}
