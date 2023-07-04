package main

import (
	"log"
	"net/http"
)

func main() {
	app := &Config{PORT, ADDR}

	server := &http.Server{
		Addr:    app.ADDR,
		Handler: app.router(),
	}

	log.Printf("Server listening on %s", app.ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
