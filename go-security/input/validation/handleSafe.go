package main

import (
	"encoding/json"
)

type PaymentSafe struct {
	Time   string
	User   string
	To     string
	Amount int
}

func handleSafe(decoder *json.Decoder) (*PaymentBasic, error) {
	p := PaymentBasic{}
	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
