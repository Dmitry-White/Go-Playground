package main

import (
	"go-security/auth/authentication/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.Handle("/messages", middleware.WithAuth(messagesHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
