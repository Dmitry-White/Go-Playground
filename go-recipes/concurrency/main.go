package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("----------- Operations -----------")
	month := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)

	go csvServer()

	fmt.Printf("%+v\n", getDistance(SEQUENTIAL_STRATEGY)(month))
	fmt.Printf("%+v\n", getDistance(CONCURRENT_STRATEGY)(month))
	fmt.Println("----------------------------------")

	fmt.Println("------------ Timeouts ------------")
	fmt.Printf("%+v\n", realTimeBidding())
	fmt.Println("----------------------------------")

	fmt.Println("----------- WaitGroups -----------")
	fmt.Printf("%+v\n", patchVMs("1.0.2"))
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Once --------------")
	m := Message{
		Content: "There is nothing more deceptive than an obvious fact.",
	}
	fmt.Printf("%+v\n", m.Sign()) // 2021/05/03 20:24:45 Calculating new signature...
	fmt.Printf("%+v\n", m.Sign()) // No-op
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Pool --------------")
	fmt.Printf("%+v\n", multiMedian())
	fmt.Println("----------------------------------")

	fmt.Println("------------- Atomic -------------")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go uploadServer(&wg)

	fmt.Printf("%+v\n", uploadSize())
	wg.Wait()
	fmt.Println("----------------------------------")

	fmt.Println("------------- Errors -------------")
	fmt.Printf("%+v\n", getErrors())
	fmt.Println("----------------------------------")
}
