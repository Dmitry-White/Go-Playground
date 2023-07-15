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
	INDEX  string
	HEALTH string
}
