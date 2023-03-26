package main

import (
	"net/http"
)

func handler(resw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	resw.Write([]byte("Placeholder"))
}
