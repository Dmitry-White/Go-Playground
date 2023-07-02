package main

import (
	"fmt"
	"net/http"
)

func handleIndex(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resw, "Index")
}
