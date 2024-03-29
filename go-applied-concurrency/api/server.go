package api

import (
	"fmt"
	"log"
	"net/http"

	"go-applied-concurrency/api/handlers"
)

func Server() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := handlers.InitRoutes(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
