package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(ROUTES.HEALTH, handleHealth)
	http.HandleFunc(ROUTES.CHECK, handleCheck)
	http.HandleFunc(ROUTES.MATH, handleMath)

	log.Printf("Server listening on %s", ADDR)

	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Fatal(err)
	}
}
