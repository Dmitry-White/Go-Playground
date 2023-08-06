package main

import (
	"go-microservices/listener-service/dal"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AppConfig struct {
	PORT     int
	ADDR     string
	AMQP     *amqp.Connection
	Services Services
}

type Exchanges struct {
	LOGS string
}

type Queues struct {
	RANDOM string
}

type Topics struct {
	LOG_INFO    string
	LOG_WARNING string
	LOG_ERROR   string
}

type Events struct {
	LOG string
}

type HandlerFunc = func(messages <-chan amqp.Delivery)

type ServiceConfig struct {
	Name string
	URL  string
}

type Services struct {
	Log    ServiceConfig
	Models dal.Models
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
