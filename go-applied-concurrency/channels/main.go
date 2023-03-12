package main

import "fmt"

func main() {
	fmt.Println("----------- Unbuffered -----------")
	unbuffered()
	fmt.Println("----------------------------------")

	fmt.Println("------------ Buffered ------------")
	buffered()
	fmt.Println("----------------------------------")
}
