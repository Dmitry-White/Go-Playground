package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.ServeMux{}

	router.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		log.Fatal(err)
	}
}
