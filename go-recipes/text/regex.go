package main

import "fmt"

type Transaction struct {
	Volume int
	Symbol string
	Price  float64
}

func parseLedger(ledger []string) []Transaction {
	fmt.Println("Not Implemented")
	return []Transaction{}
}
