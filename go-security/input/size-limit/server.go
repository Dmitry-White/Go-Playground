package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	n := len(data)

	payload := fmt.Sprintf("%d bytes stored", n)
	fmt.Fprintln(res, payload)
}
