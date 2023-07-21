package main

type AppConfig struct {
	PORT int
	ADDR string
}

type Routes struct {
	INDEX   string
	HEALTH  string
	PROCESS string
}

type ServiceConfig struct {
	Name string
	URL  string
}

type Services struct {
	Auth ServiceConfig
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}
