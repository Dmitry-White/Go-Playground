package main

import (
	"fmt"
	"time"
)

func selected() {
	// Create 2 buffered channels
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// start the 2 greeters to provide sayings
	go greetSelected(greetings, ch1)
	go greetSelected(goodbyes, ch2)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	for {
		// Receive greetings
		select {
		case greeting := <-ch1:
			// Sleep and print
			printSaying(greeting)
		case goodbye := <-ch2:
			// Sleep and print
			printSaying(goodbye)
		default:
			return
		}
	}
}
