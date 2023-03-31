package main

import (
	"fmt"
	"net/http"
)

const (
	certFile = "cert.pem"
	keyFile  = "key.pem"
)

func healthHandler(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resw, "OK\n")
}

func main() {
	router := http.ServeMux{}

	router.HandleFunc("/health", healthHandler)

	http.ListenAndServeTLS(":8080", certFile, keyFile, &router)
}
