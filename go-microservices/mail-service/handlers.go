package main

import (
	"log"
	"net/http"
)

func (app *AppConfig) handleMail(resw http.ResponseWriter, req *http.Request) {
	log.Println("Mail sent!")
}
