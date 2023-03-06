package main

import (
	"go-rest-api/api"
	"go-rest-api/api/data"
)

func main() {
	app := api.App{
		App: &data.App{
			Port: ":8000",
		},
	}
	app.Init()
	app.Run()
}
