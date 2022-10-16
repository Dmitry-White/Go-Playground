package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	adUrl string
	price int
}

var URLS = map[string]string{
	"CONTINUE":  "https://http.cat/100",
	"TEAPOT":    "https://http.cat/418",
	"NOT_FOUND": "https://http.cat/404",
}

var defaultBid = Bid{
	adUrl: URLS["CONTINUE"],
	price: 100,
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	resultingBid := Bid{
		adUrl: url,
		price: 20,
	}

	return resultingBid
}

func findBid(ctx context.Context, url string) Bid {
	fmt.Println("Processing...")

	// Buffered to avoid Goroutine leak,
	// meaning Main Goroutine will hang
	// when there's a Child Goroutine waiting to be read
	channel := make(chan Bid, 1)

	go func() {
		channel <- bestBid(url)
	}()

	select {
	case bid := <-channel:
		fmt.Println("Custom Bid Done")
		return bid
	case <-ctx.Done():
		fmt.Println("Custom Bid Timeout")
		return defaultBid
	}

}

func timedBid(url string, timeout time.Duration) string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	bid := findBid(ctx, url)

	return fmt.Sprintf("Bid: %v", bid)
}

func main() {
	fmt.Println(timedBid(URLS["TEAPOT"], 100*time.Millisecond))
	fmt.Println("-------------")

	fmt.Println(timedBid(URLS["NOT_FOUND"], 10*time.Millisecond))
	fmt.Println("-------------")
}
