package main

import "go-rest-api/api"

func main() {
	app := api.App{
		Port: ":8000",
	}
	app.Init()
	app.Run()
}
