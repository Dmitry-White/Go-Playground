package main

type AppConfig struct {
	PORT     int
	ADDR     string
	Services Services
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
	Log  ServiceConfig
	Mail ServiceConfig
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

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}
