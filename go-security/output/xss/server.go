package main

import (
	"go-security/output/xss/dal"
	"log"
	"net/http"
)

func handler(resw http.ResponseWriter, req *http.Request) {
	user := req.URL.Path[1:] // "/frodo" -> "frodo"

	ms, err := dal.LoadMessages(user)
	if err != nil {
		http.Error(resw, "can't get messages", http.StatusBadRequest)
		return
	}

	renderFunc, err := getRenderFunc(SAFE_STRATEGY)
	if err != nil {
		log.Fatalln(err)
	}

	err = renderFunc(resw, ms, user)
	if err != nil {
		http.Error(resw, "can't render messages", http.StatusInternalServerError)
		return
	}
}
