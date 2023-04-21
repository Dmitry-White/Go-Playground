package main

import (
	"regexp"
	"strconv"
)

type Transaction struct {
	Volume int
	Symbol string
	Price  float64
}

/*
Examples:

12 shares of MSFT for $234.57

10 shares of TSLA for $692.4
*/
var transactionRegex = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)

func parseLedger(ledger []string) []Transaction {
	transactions := []Transaction{}

	for _, input := range ledger {
		matches := transactionRegex.FindStringSubmatch(input)

		volume, _ := strconv.Atoi(matches[1])
		symbol := matches[2]
		price, _ := strconv.ParseFloat(matches[3], 64)

		transaction := Transaction{volume, symbol, price}

		transactions = append(transactions, transaction)
	}

	return transactions
}
