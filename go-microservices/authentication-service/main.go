package main

import (
	"go-microservices/authentication-service/dal"
	"log"
	"net/http"
)

func main() {
	connection := dal.Connect()
	if connection == nil {
		log.Fatalln("Can't connect to DB!")
	}

	log.Printf("DB Stats: %+v\n", connection.Stats())

	app := &AppConfig{
		PORT: PORT,
		ADDR: ADDR,
		DB:   connection,
		Services: Services{
			Models: dal.New(connection),
			Broker: SERVICES.Broker,
		},
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
