package main

import (
	"go-microservices/logger-service/dal"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	PORT     int
	PORT_RPC int
	ADDR     string
	ADDR_RPC string
	DB       *mongo.Client
	Services Services
}

type Routes struct {
	LOG    string
	HEALTH string
}

type Services struct {
	Models dal.Models
}

type RequestPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogRPCPayload struct {
	RequestPayload
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
