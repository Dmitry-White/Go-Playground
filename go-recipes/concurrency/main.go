package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------- Operations -----------")
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%+v\n", operations(SEQUENTIAL_STRATEGY)(month))
	fmt.Printf("%+v\n", operations(CONCURRENT_STRATEGY)(month))
	fmt.Println("----------------------------------")
}
