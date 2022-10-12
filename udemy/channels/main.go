package main

import (
	"fmt"
	"net/http"
)

func checkLink(channel chan string, link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return false
	}
	fmt.Println(link, "is up!")
	channel <- link
	return true
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(channel, link)
	}

	// Naive loop to wait for a message from a channel
	// number of time a goroutine was created
	// for i := 0; i < len(links); i++ {
	// 	fmt.Printf("Status is %v\n", <-channel)
	// }

	// Infinite loop with Message Syntax
	// to kick off a new goroutine each time
	// a channel receives a message from the completed goroutine
	// for {
	// 	go checkLink(channel, <-channel)
	// }

	// Infinite loop with Range Syntax
	// to kick off a new goroutine each time
	// a channel receives a message from the completed goroutine
	for link := range channel {
		go checkLink(channel, link)
	}
}
