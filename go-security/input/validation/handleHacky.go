package main

import (
	"encoding/json"
)

type PaymentHacky struct {
	Time   string
	User   string
	To     string
	Amount int
}

func handleHacky(decoder *json.Decoder) (*PaymentBasic, error) {
	p := PaymentBasic{}
	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
