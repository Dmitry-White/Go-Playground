package main

import (
	"fmt"
	"os"
)

const PORT = 80
const READ_LIMIT = 1048576

var ROUTES = Routes{
	INDEX:   "/",
	HEALTH:  "/health",
	PROCESS: "/process",
}

var EXCHANGES = &Exchanges{
	LOGS: "logs_topic",
}

var TOPICS = &Topics{
	LOG_INFO:    "log.INFO",
	LOG_WARNING: "log.WARNING",
	LOG_ERROR:   "log.ERROR",
}

var SERVICES = Services{
	Auth: ServiceConfig{
		Name: "Auth",
		URL:  "http://authentication-service/auth",
	},
	Log: ServiceConfig{
		Name: "Log",
		URL:  "http://logger-service/log",
	},
	Mail: ServiceConfig{
		Name: "Mail",
		URL:  "http://mail-service/mail",
	},
	Async: ServiceConfig{
		Name: "Async",
		URL:  os.Getenv("DSN"),
	},
}

var ADDR = fmt.Sprintf(":%d", PORT)
