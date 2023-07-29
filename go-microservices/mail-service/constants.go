package main

import "fmt"

const PORT = 80

var ROUTES = &Routes{
	MAIL:   "/mail",
	HEALTH: "/health",
}

var ADDR = fmt.Sprintf(":%d", PORT)
