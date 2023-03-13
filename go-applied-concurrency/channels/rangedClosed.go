package main

import (
	"fmt"
	"time"
)

func rangedClosed() {
	// Create a bidirectional channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greetRangeClosed(ch)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	// Main goroutine will end up forever printing empty line
	// since receiving from a closed channel
	// will immediately complete with a zero value
	// of the channel type
	for {
		// Receive greetings
		greeting := <-ch

		// Sleep and print
		time.Sleep(time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}
