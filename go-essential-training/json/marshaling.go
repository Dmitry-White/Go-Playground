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

func marshalResponse(amount float64) {
	// Loaded from a DB
	const currentBalance = 10000.0

	// Marshal response
	response := Response{
		Ok:      true,
		Balance: currentBalance + amount,
	}

	responseBytes, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("error: can't marshal - %s", err)
	}

	log.Printf("Response: %+v\n", string(responseBytes))
}
