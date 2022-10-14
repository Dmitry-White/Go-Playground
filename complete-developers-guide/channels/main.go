package main

import (
	"fmt"
	"net/http"
	"time"
)

// Runs inside a Child Goroutine
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

// Spawns the Main Goroutine
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Receives messages from any Child Goroutine
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
	// for link := range channel {
	// 	go checkLink(channel, link)
	// }

	// Timed Infinite loop with Range Syntax
	// to kick off a new goroutine each time
	// a channel receives a message from the completed goroutine.
	// Each time any Child Goroutine sends a message
	// through the channel back to Main Goroutine
	// "link" variable, that points to the same place in memory,
	// will be updated and thus that same place in memory
	// will hold a different value.
	// That may become a problem if any Child Goroutine relies on it
	// while it's still runnning
	for link := range channel {
		// Since Go passes stuff by value,
		// "linkCopy" will be a new place in memory
		// to hold a copy of the link for Child Goroutine to work on
		// without worrying about Main Goroutine updating it
		go func(linkCopy string) {
			// Pauses Goroutine to avoid flooding the links,
			// it's placed outside to keep checkLink function's
			// single responsibility
			time.Sleep(5 * time.Second)
			checkLink(channel, linkCopy)
		}(link)
	}
}
