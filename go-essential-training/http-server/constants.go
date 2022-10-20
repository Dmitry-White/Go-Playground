package main

import "fmt"

const PORT = 8080

type Routes struct {
	HEALTH string
	MATH   string
}

var ROUTES = Routes{
	HEALTH: "/health",
	MATH:   "/math",
}

var ADDR = fmt.Sprintf("localhost:%d", PORT)
