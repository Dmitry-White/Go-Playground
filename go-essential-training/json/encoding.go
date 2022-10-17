package main

import (
	"encoding/json"
	"log"
	"strings"
)

func decodeRequest() {
	// Simulate a file/socket
	requestReader := strings.NewReader(DATA)

	// Create request decoder
	requestDecoder := json.NewDecoder(requestReader)

	// Decode request
	var request Request
	err := requestDecoder.Decode(&request)

	if err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	log.Printf("Request: %+v\n", request)
}
