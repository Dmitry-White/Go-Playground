package main

import "fmt"

const PORT = 8080

var ROUTES = Routes{
	HEALTH: "/health",
	CHECK:  "/check",
	MATH:   "/math",
}

var ADDR = fmt.Sprintf("localhost:%d", PORT)
