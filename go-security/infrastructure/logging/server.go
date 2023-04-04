package main

import (
	"expvar"
	"log"
	"net/http"
)

var loginCalls = expvar.NewInt("login.calls")

func loginHandler(resw http.ResponseWriter, req *http.Request) {
	log.Println("Not Implemented")
	loginCalls.Add(1)
}
