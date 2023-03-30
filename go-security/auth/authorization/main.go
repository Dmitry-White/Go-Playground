package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.ServeMux{}

	router.HandleFunc("/profile", handler)

	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		log.Fatalln(err)
	}
}
