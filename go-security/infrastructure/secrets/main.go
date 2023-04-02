package main

import (
	"log"
	"net/http"
)

// init is automatically invoked before main()
func init() {
	prepareEnv()
}

func main() {
	router := http.ServeMux{}

	router.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		log.Fatalln(err)
	}
}
