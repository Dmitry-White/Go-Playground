package main

import (
	"log"
	"net/http"
)

func handleLog(resw http.ResponseWriter, req *http.Request) {
	log.Println("Logged!")
}
