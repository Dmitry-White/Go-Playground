package main

import (
	"encoding/json"
	"fmt"
)

// LineItem is a line in receipt
type LineItem struct {
	SKU      string
	Price    float64
	Discount float64
	Quantity int
}

func parseLineItem(data []byte) (LineItem, error) {
	// Defaulting a field to non-zero so that
	// it's clear when it's overriden by a zero value
	li := LineItem{
		Quantity: 1,
	}

	err := json.Unmarshal(data, &li)
	if err != nil {
		return LineItem{}, err
	}

	if li.Quantity < 1 {
		return LineItem{}, fmt.Errorf("bad quantity")
	}

	return li, nil
}

func handleZero() interface{} {
	firstLine := []byte(`{"sku": "x3xs", "price": 1.2}`)
	firstLineItem, err := parseLineItem(firstLine)
	if err != nil {
		fmt.Println("First line:", err)
	}

	secondLine := []byte(`{"sku": "x3xs", "price": 1.2, "quantity": 0}`)
	secondLineItem, err := parseLineItem(secondLine)
	if err != nil {
		fmt.Println("Second line:", err)
	}

	return []LineItem{firstLineItem, secondLineItem}
}
