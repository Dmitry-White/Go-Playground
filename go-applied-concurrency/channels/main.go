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

	fmt.Println("------------- Closed -------------")
	closed()
	fmt.Println("----------------------------------")
}
