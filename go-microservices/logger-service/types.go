package main

type AppConfig struct {
	PORT int
	ADDR string
}

type Routes struct {
	LOG    string
	HEALTH string
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
