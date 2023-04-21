package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("------------- Format -------------")
	trades := []Trade{
		{"MSFT", 231, 234.57},
		{"TSLA", 123, 686.75},
		{"BRK-A", 1, 399100},
	}
	genReport(os.Stdout, trades)
	fmt.Println("----------------------------------")

	fmt.Println("------------ Unicode -------------")
	fmt.Println(sentenceLength([]string{"«", "Don't", "Panic", "»"}))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Cases --------------")
	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
	fmt.Println("----------------------------------")
}
