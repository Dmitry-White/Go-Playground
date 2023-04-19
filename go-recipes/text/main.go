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
}
