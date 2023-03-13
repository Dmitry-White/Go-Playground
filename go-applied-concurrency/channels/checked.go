package main

import (
	"fmt"
	"time"
)

func checked() {
	// Create a bidirectional channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greetRangeClosed(ch)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	// Channels have a optional second return value
	// to check whether a received value came from a closed channel or not
	// If it's false, then the channel is closed.
	for {
		// Receive greetings
		greeting, ok := <-ch
		if !ok {
			return
		}

		// Sleep and print
		time.Sleep(time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}
