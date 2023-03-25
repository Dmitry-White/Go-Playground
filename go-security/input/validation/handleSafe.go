package main

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type PaymentSafe struct {
	Time   string `json:"time"`
	User   string `json:"user" validate:"required,min=2,max=100"`
	To     string `json:"to"`
	Amount int    `json:"amount" validate:"required,min=2,max=10000"`
}

func (p *PaymentSafe) Validate() error {
	tagValidator := validator.New()

	err := tagValidator.Struct(p)
	if err != nil {
		return err
	}

	return nil
}

func handleSafe(decoder *json.Decoder) (Payment, error) {
	p := PaymentSafe{}

	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}

	err = p.Validate()
	if err != nil {
		return nil, err
	}

	return &p, nil
}
