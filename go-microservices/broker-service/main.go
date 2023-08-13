package main

import (
	"go-microservices/broker-service/dal"
	"log"
	"net/http"
)

func main() {
	connection := dal.Connect()
	if connection == nil {
		log.Fatalln("Can't connect to Queue!")
	}
	defer connection.Close()

	log.Printf("Queue Config: %+v\n", connection.Config)

	app := &AppConfig{
		PORT: PORT,
		ADDR: ADDR,
		AMQP: connection,
		Services: Services{
			Auth:    SERVICES.Auth,
			Log:     SERVICES.Log,
			Mail:    SERVICES.Mail,
			Async:   SERVICES.Async,
			LogRPC:  SERVICES.LogRPC,
			LogGRPC: SERVICES.LogGRPC,
			Models:  dal.New(connection),
		},
	}

	handlerErr := app.handleSetup()
	if handlerErr != nil {
		log.Fatalln("Can't setup Queue!", handlerErr)
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
