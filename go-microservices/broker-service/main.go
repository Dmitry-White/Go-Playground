package main

import (
	"log"
	"net/http"
)

func main() {
	app := &AppConfig{
		PORT: PORT,
		ADDR: ADDR,
		Services: Services{
			Auth: SERVICES.Auth,
			Log:  SERVICES.Log,
			Mail: SERVICES.Mail,
		},
	}

	server := &http.Server{
		Addr:    app.ADDR,
		Handler: app.router(),
	}

	log.Printf("Server listening on %s", app.ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
