package main

import "fmt"

func main() {
	fmt.Println("------------ Unmarshal -----------")
	fmt.Printf("%+v\n", getTemperature())
	fmt.Println("----------------------------------")

	fmt.Println("------------- Complex ------------")
	fmt.Printf("%+v\n", scanStations())
	fmt.Println("----------------------------------")
}
