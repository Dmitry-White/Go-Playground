package main

import "fmt"

const PORT = 80
const READ_LIMIT = 1048576

var ROUTES = Routes{
	INDEX:   "/",
	HEALTH:  "/health",
	PROCESS: "/process",
}

var SERVICES = Services{
	Auth: ServiceConfig{
		Name: "Auth",
		URL:  "http://authentication-service/auth",
	},
}

var ADDR = fmt.Sprintf(":%d", PORT)
