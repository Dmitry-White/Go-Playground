package main

import (
	"go-microservices/logger-service/dal"
	"log"
	"net/http"
)

func main() {
	connection := dal.Connect()
	if connection == nil {
		log.Fatalln("Can't connect to DB!")
	}

	log.Printf("DB Sessions: %+v\n", connection.NumberSessionsInProgress())

	app := &AppConfig{PORT, ADDR}

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
