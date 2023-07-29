package main

import (
	"log"
	"net/http"
)

func main() {
	app := &AppConfig{PORT, ADDR}

	server := http.Server{
		Addr: app.ADDR,
	}

	log.Printf("Server listening on %s", app.ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
