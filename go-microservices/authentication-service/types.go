package main

import (
	"database/sql"
	"go-microservices/authentication-service/dal"
)

type AppConfig struct {
	PORT     int
	ADDR     string
	DB       *sql.DB
	Services Services
}

type Routes struct {
	AUTH   string
	HEALTH string
}

type ServiceConfig struct {
	Name string
	URL  string
}

type Services struct {
	Broker ServiceConfig
	Models dal.Models
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type RequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BrokerRequestPayload struct {
	Action string     `json:"action"`
	Log    LogPayload `json:"log,omitempty"`
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
