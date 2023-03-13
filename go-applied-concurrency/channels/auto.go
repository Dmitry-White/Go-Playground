package main

import (
	"fmt"
	"time"
)

func auto() {
	// Create a bidirectional channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greetRangeClosed(ch)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	// Channels support for-range operator
	// to iterate over the values received from the channel.
	// The range automatically exists when the channel closes.

	// Receive greetings
	for greeting := range ch {
		// Sleep and print
		time.Sleep(time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}
