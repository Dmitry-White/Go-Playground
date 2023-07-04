package main

import (
	"encoding/json"
)

func generateJson() ([]byte, error) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	data, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return nil, err
	}

	return data, nil
}
