package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------- Struct -------------")
	fmt.Printf("%+v\n", getDoorEvent("front door", time.Now(), "open"))
	fmt.Printf("%+v\n", getTermostatEvent("front door", time.Now(), 25))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Methods ------------")
	fmt.Printf("%+v\n", getThermostat())
	fmt.Println("----------------------------------")

	fmt.Println("----------- Interfaces -----------")
	fmt.Printf("%+v\n", getSensors())
	fmt.Println("----------------------------------")
}
