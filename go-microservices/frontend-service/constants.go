package main

import "fmt"

const PORT = 80

var PATHS = Paths{
	JS:        "./web/js",
	PAGES:     "./web/pages",
	TEMPLATES: "./web/templates",
}

var ROUTES = Routes{
	INDEX: "/",
	JS:    "/js/",
}

var ADDR = fmt.Sprintf(":%d", PORT)
