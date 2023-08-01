package main

import (
	"go-microservices/mail-service/smtp"

	mail "github.com/xhit/go-simple-mail/v2"
)

type AppConfig struct {
	PORT     int
	ADDR     string
	SMTP     *mail.SMTPClient
	Services Services
}

type Routes struct {
	MAIL   string
	HEALTH string
}

type Services struct {
	Models smtp.Models
}

type RequestPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
