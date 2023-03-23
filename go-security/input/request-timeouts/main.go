package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	server, err := getClient("safe")()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(server)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
