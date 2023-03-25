package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type PaymentBasic struct {
	Time   string
	User   string
	To     string
	Amount int
}

func (p *PaymentBasic) validate() error {
	foundErrors := []error{}

	if len(p.User) == 0 {
		log.Printf("Format with +: %+v\n", p)

		err := fmt.Errorf("bad User in %#v", p)
		foundErrors = append(foundErrors, err)
	}

	if p.Amount <= 0 {
		log.Printf("Format with +: %+v\n", p)

		err := fmt.Errorf("bad Amount in %#v", p)
		foundErrors = append(foundErrors, err)
	}

	if len(foundErrors) > 0 {
		return errors.Join(foundErrors...)
	}

	return nil
}

func handleBasic(decoder *json.Decoder) (*PaymentBasic, error) {
	p := PaymentBasic{}
	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}

	err = p.validate()
	if err != nil {
		return nil, err
	}

	return &p, nil
}
