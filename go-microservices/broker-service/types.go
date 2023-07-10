package main

type AppConfig struct {
	PORT int
	ADDR string
}

type Routes struct {
	INDEX  string
	HEALTH string
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
