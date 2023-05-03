package main

import "fmt"

func main() {
	fmt.Println("------------- Client -------------")
	fmt.Printf("%+v\n", sendMetrics())
	fmt.Println("----------------------------------")
}
