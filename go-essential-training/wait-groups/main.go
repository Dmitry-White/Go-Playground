package main

import (
	"fmt"
	"time"
)

func fetchSerial(urls []string) {
	fmt.Println("----------------")
	fmt.Println("Fetch Serial")
	fmt.Println(urls)
	fmt.Println("----------------")
}

func fetchConcurrent(urls []string) {
	fmt.Println("----------------")
	fmt.Println("Fetch Concurrent")
	fmt.Println(urls)
	fmt.Println("----------------")
}

func main() {
	fmt.Println("Wait Groups")

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://google.com",
		"https://amazon.com",
	}

	startSerial := time.Now()
	fetchSerial(urls)
	endSerial := time.Since(startSerial)
	fmt.Println("Time Elapsed: ", endSerial)

	startConcurrent := time.Now()
	fetchConcurrent(urls)
	endConcurrent := time.Since(startConcurrent)
	fmt.Println("Time Elapsed: ", endConcurrent)
}
