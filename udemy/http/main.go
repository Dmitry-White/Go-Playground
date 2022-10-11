package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", *resp)
}
