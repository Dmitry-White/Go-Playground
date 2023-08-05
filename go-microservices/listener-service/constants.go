package main

import "fmt"

const PORT = 80

var EXCHANGES = &Exchanges{
	LOGS: "logs_topic",
}

var QUEUES = &Queues{
	RANDOM: "",
}

var TOPICS = &Topics{
	LOG_INFO:    "log.INFO",
	LOG_WARNING: "log.WARNING",
	LOG_ERROR:   "log.ERROR",
}

var SERVICES = Services{
	Log: ServiceConfig{
		Name: "Log",
		URL:  "http://logger-service/log",
	},
}

var ADDR = fmt.Sprintf(":%d", PORT)
