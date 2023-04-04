package main

import (
	"expvar"
	"log"
	"net/http"
)

/*
	`expvar` is the native language library if we want to instrument or add metrics to a Go app.
	When we import the expvar library, it's `init` registers http.Handler and a couple of variables.
	This registers `/debug/vars` at default http ServeMux as an endpoint
	where all the metrics will be published in JSON format.
	Using your own ServeMux requires explicit `/debug/vars` registering.
*/

func main() {
	router := http.ServeMux{}

	router.HandleFunc("/login", loginHandler)
	router.Handle("/debug/vars", expvar.Handler())

	err := http.ListenAndServe(":8080", &router)
	if err != nil {
		log.Fatal(err)
	}
}
