package main

import (
	"go-security/auth/authorization/middleware"
	"log"
	"net/http"
)

func main() {
	router := http.ServeMux{}

	router.Handle("/profile", middleware.WithAccess(handler))

	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		log.Fatalln(err)
	}
}
