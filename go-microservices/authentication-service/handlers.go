package main

import (
	"log"
	"net/http"
)

func (a *AppConfig) handleIndex(resw http.ResponseWriter, req *http.Request) {
	log.Println("Authenticated")
}
