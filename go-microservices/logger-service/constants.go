package main

import "fmt"

const PORT = 80
const READ_LIMIT = 1048576

var ROUTES = &Routes{
	LOG:    "/log",
	HEALTH: "/health",
}

var ADDR = fmt.Sprintf(":%d", PORT)
