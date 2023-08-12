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

	app := &AppConfig{
		PORT:      PORT,
		PORT_RPC:  PORT_RPC,
		ADDR:      ADDR,
		ADDR_RPC:  ADDR_RPC,
		ADDR_GRPC: ADDR_GRPC,
		DB:        connection,
		Services: Services{
			Models: dal.New(connection),
		},
	}

	handlerErr := app.handleSetup()
	if handlerErr != nil {
		log.Fatalln("Can't setup RPC!", handlerErr)
	}

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
