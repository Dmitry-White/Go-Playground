package main

import (
	"go-microservices/broker-service/dal"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AppConfig struct {
	PORT     int
	ADDR     string
	AMQP     *amqp.Connection
	Services Services
}

type Routes struct {
	INDEX   string
	HEALTH  string
	PROCESS string
}

type Exchanges struct {
	LOGS string
}

type Topics struct {
	LOG_INFO    string
	LOG_WARNING string
	LOG_ERROR   string
}

type ServiceConfig struct {
	Name string
	URL  string
}

type Services struct {
	Auth    ServiceConfig
	Log     ServiceConfig
	Mail    ServiceConfig
	Async   ServiceConfig
	LogRPC  ServiceConfig
	LogGRPC ServiceConfig
	Models  dal.Models
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type AsyncPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogRPCPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogGRPCPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type RequestPayload struct {
	Action  string         `json:"action"`
	Auth    AuthPayload    `json:"auth,omitempty"`
	Log     LogPayload     `json:"log,omitempty"`
	Mail    MailPayload    `json:"mail,omitempty"`
	Async   AsyncPayload   `json:"async,omitempty"`
	LogRPC  LogRPCPayload  `json:"logRPC,omitempty"`
	LogGRPC LogGRPCPayload `json:"logGRPC,omitempty"`
}

type RequestPayloadRPC struct {
	Name string
	Data string
}
