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

	log.Printf("Queue Config: %+v\n", connection.Config)
}
