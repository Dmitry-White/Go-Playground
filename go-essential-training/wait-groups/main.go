package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchContentType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%s -> Error: %v", url, err)
	}

	contentType := resp.Header.Get("Content-Type")
	return fmt.Sprintf("%s -> %s\n", url, contentType)
}

func fetchSerial(urls []string) {
	fmt.Println("----------------")
	fmt.Println("Fetch Serial")

	for _, url := range urls {
		fmt.Println(fetchContentType(url))
	}

	fmt.Println("----------------")
}

func fetchConcurrent(urls []string) {
	fmt.Println("----------------")
	fmt.Println("Fetch Concurrent")

	var waitGroup sync.WaitGroup

	for _, url := range urls {
		waitGroup.Add(1)
		go func(url string) {
			fmt.Println(fetchContentType(url))
			waitGroup.Done()
		}(url)
	}

	waitGroup.Wait()

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
