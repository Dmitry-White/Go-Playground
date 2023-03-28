package main

import (
	"go-security/output/sensitive-data/dal"
	"log"
	"net/http"
)

func handler(resw http.ResponseWriter, req *http.Request) {
	login := req.URL.Query().Get("user")
	if login == "nil" {
		http.Error(resw, "Login query was not found", http.StatusBadRequest)
		return
	}

	user := dal.GetUser(login)
	if user == nil {
		http.Error(resw, "No user with this login", http.StatusNotFound)
		return
	}

	handleData, err := getDataHandler(SAFE_STRATEGY)
	if err != nil {
		log.Fatalln(err)
	}

	handleData(resw, user)
}
