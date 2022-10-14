package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchContentType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n", url, contentType)
}

func fetchSerial(urls []string) {
	fmt.Println("----------------")
	fmt.Println("Fetch Serial")

	for _, url := range urls {
		fetchContentType(url)
	}

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
		"https://go.dev",
		"https://api.github.com",
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://nodejs.org",
		"https://go.dev",
		"https://api.github.com",
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://nodejs.org",
		"https://go.dev",
		"https://api.github.com",
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://nodejs.org",
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
