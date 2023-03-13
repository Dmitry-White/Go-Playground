package main

import "fmt"

func main() {
	fmt.Println("----------- Unbuffered -----------")
	unbuffered()
	fmt.Println("----------------------------------")

	fmt.Println("------------ Buffered ------------")
	buffered()
	fmt.Println("----------------------------------")

	fmt.Println("--------- Unidirectional ---------")
	unidirectional()
	fmt.Println("----------------------------------")

	// fmt.Println("------------- Ranged -------------")
	// ranged()
	// fmt.Println("----------------------------------")

	// fmt.Println("--------- Ranged Closed ----------")
	// rangedClosed()
	// fmt.Println("----------------------------------")

	fmt.Println("------------- Checked -------------")
	checked()
	fmt.Println("----------------------------------")

	fmt.Println("------------- Auto -------------")
	auto()
	fmt.Println("----------------------------------")

	// fmt.Println("------------ Selected ------------")
	// selected()
	// fmt.Println("----------------------------------")

	fmt.Println("--------- Selected Checked --------")
	selectedChecked()
	fmt.Println("----------------------------------")
}
