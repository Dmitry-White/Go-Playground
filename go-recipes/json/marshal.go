package main

import (
	"encoding/json"
	"os"
)

// Quantity is combination of value and unit (e.g. 2.7cm)
type Quantity struct {
	Value float64
	Unit  string
}

func getQuantity() interface{} {
	q := Quantity{1.78, "meter"}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(&q)
	if err != nil {
		return err
	}

	return &q
}
