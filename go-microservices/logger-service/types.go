package main

import (
	"go-microservices/logger-service/dal"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	PORT   int
	ADDR   string
	DB     *mongo.Client
	Models dal.Models
}

type Routes struct {
	LOG    string
	HEALTH string
}

type RequestPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
