package main

import (
	"fmt"
	"net/http"
)

func handleMetric(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resw, "OK\n")
}

func startServer() interface{} {
	router := http.ServeMux{}

	router.HandleFunc("/metric", handleMetric)

	err := http.ListenAndServe(":8000", &router)
	if err != nil {
		return nil
	}
	return &router
}
