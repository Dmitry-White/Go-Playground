package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("----------- Operations -----------")
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go server(&wg)

	fmt.Printf("%+v\n", operations(SEQUENTIAL_STRATEGY)(month))
	fmt.Printf("%+v\n", operations(CONCURRENT_STRATEGY)(month))

	wg.Wait()
	fmt.Println("----------------------------------")
}
