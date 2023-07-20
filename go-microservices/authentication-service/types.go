package main

import (
	"database/sql"
	"go-microservices/authentication-service/dal"
)

type AppConfig struct {
	PORT   int
	ADDR   string
	DB     *sql.DB
	Models dal.Models
}

type Routes struct {
	AUTH   string
	HEALTH string
}

type RequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
