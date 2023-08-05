package main

import (
	"go-microservices/mail-service/smtp"
	"log"
	"net/http"
)

func main() {
	connection, mailer := smtp.Connect()
	if connection == nil {
		log.Fatalln("Can't connect to SMTP Server!")
	}

	log.Printf("SMTP Client: %+v\n", connection.Client)

	app := &AppConfig{
		PORT: PORT,
		ADDR: ADDR,
		SMTP: connection,
		Services: Services{
			Models: smtp.New(connection, mailer),
		},
	}

	server := http.Server{
		Addr:    app.ADDR,
		Handler: app.router(),
	}

	log.Printf("Server listening on %s", app.ADDR)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
