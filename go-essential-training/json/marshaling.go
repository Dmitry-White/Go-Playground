package main

import (
	"encoding/json"
	"log"
)

func unmarshalRequest() {
	// Simulate binary data
	requestBytes := []byte(DATA)

	// Unmarshal request
	var request Request
	err := json.Unmarshal(requestBytes, &request)

	if err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	log.Printf("Request: %+v\n", request)
}
