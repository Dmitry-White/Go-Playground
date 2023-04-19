package main

import (
	"fmt"
	"io"
	"log"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
}

func genReport(w io.Writer, trades []Trade) {
	for i, trade := range trades {
		log.Printf("%d: %#v", i, trade)
		// ... 2: main.Trade{Symbol:"BRK-A", Volume:1, Price:399100}
		fmt.Fprintf(w, "%-10s %04d %.2f\n", trade.Symbol, trade.Volume, trade.Price)
		// MSFT       0231 234.57
	}
}
