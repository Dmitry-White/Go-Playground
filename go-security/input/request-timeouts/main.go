package main

import (
	"log"
)

func main() {
	server, err := getClient("safe")()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(server)
}
