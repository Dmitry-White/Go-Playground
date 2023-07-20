package main

import "fmt"

const PORT = 80

var ROUTES = Routes{
	INDEX: "/",
	JS:    "/js/",
}

var PATHS = Paths{
	JS:        "./web/js",
	PAGES:     "./web/pages",
	TEMPLATES: "./web/templates",
}

var ADDR = fmt.Sprintf(":%d", PORT)
