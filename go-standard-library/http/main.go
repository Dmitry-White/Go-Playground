package main

import "fmt"

func main() {
	fmt.Println("------------ URLs ------------")
	urls()
	fmt.Println("------------------------------")

	fmt.Println("------------ GET -------------")
	get()
	fmt.Println("------------------------------")

	fmt.Println("------------ POST ------------")
	post()
	fmt.Println("------------------------------")

	fmt.Println("-------- JSON Encoding -------")
	jsonEncoding()
	fmt.Println("------------------------------")

	fmt.Println("-------- JSON Decoding -------")
	jsonDecoding()
	fmt.Println("------------------------------")
}
