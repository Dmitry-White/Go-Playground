package main

import (
	"encoding/json"
	"log"
	"os"
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

func encodeResponse(amount float64) {
	// Loaded from a DB
	const currentBalance = 10000.0

	// Simulate a file/socket
	responseWriter := os.Stdout

	// Create response encoder
	responseEncoder := json.NewEncoder(responseWriter)

	// Encode request
	response := Response{
		Ok:      true,
		Balance: currentBalance + amount,
	}
	err := responseEncoder.Encode(response)

	if err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}
