package main

import (
	"go-microservices/listener-service/dal"
	"log"
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
			Models: dal.New(connection),
		},
	}

	handlerErr := app.handleTopics()
	if handlerErr != nil {
		log.Fatalln("Can't handle Queue!", handlerErr)
	}

	log.Printf("Server listening on %s", app.ADDR)

	err := app.listen()
	if err != nil {
		log.Panic(err)
	}
}
