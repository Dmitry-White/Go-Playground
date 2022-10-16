package main

import (
	"context"
	"fmt"
	"time"
)

var URLS = map[string]string{
	"TEAPOT":    "https://http.cat/418",
	"NOT_FOUND": "https://http.cat/404",
}

func findBid(ctx context.Context, url string) string {
	fmt.Println("Processing...")
	fmt.Println(url)
	fmt.Println("Done")
	return url
}

func timedBid(url string, timeout time.Duration) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	bid := findBid(ctx, url)

	return fmt.Sprintf("Bid: %v", bid)
}

func main() {
	fmt.Println(timedBid(URLS["TEAPOT"], 100*time.Microsecond))
	fmt.Println("-------------")

	fmt.Println(timedBid(URLS["NOT_FOUND"], 10*time.Microsecond))
	fmt.Println("-------------")
}
