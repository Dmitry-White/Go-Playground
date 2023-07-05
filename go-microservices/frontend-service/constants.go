package main

import "fmt"

const PORT = 80

var ROUTES = Routes{
	INDEX: "/",
}

var ADDR = fmt.Sprintf(":%d", PORT)
