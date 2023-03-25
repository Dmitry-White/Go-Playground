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

func (p *PaymentHacky) Validate() error {
	return nil
}

func handleHacky(decoder *json.Decoder) (Payment, error) {
	p := PaymentHacky{}

	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
