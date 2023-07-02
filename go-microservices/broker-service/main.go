package main

import (
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ADDR,
		Handler: http.HandlerFunc(handleIndex),
	}

	log.Printf("Server listening on %s", ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
