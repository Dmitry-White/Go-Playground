package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("------------- Client -------------")
	fmt.Printf("%+v\n", sendMetrics())
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Auth --------------")
	fmt.Printf("%+v\n", authenticate())
	fmt.Println("----------------------------------")

	wg := sync.WaitGroup{}

	fmt.Println("------------- Server -------------")
	wg.Add(1)
	go startServer(&wg)
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Mux ---------------")
	wg.Add(1)
	go startServerMux(&wg)
	fmt.Println("----------------------------------")

	wg.Wait()
}
