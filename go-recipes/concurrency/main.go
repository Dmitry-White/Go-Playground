package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------- Operations -----------")
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)

	go server()

	fmt.Printf("%+v\n", getDistance(SEQUENTIAL_STRATEGY)(month))
	fmt.Printf("%+v\n", getDistance(CONCURRENT_STRATEGY)(month))
	fmt.Println("----------------------------------")

	fmt.Println("------------ Timeouts ------------")
	fmt.Printf("%+v\n", realTimeBidding())
	fmt.Println("----------------------------------")
}
