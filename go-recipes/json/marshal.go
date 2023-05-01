package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Quantity is combination of value and unit (e.g. 2.7cm)
type Quantity struct {
	Value float64
	Unit  string
}

// MarshalJSON implements the json.Marshaler interface
func (q *Quantity) MarshalJSON() ([]byte, error) {
	if q.Unit == "" {
		return nil, fmt.Errorf("empty unit")
	}

	data := fmt.Sprintf("%f%s", q.Value, q.Unit)

	return json.Marshal(data)
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
