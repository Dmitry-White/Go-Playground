package main

import (
	"fmt"
	"time"
)

func selectedChecked() {
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
		case greeting, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			// Sleep and print
			printSaying(greeting)
		case goodbye, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			// Sleep and print
			printSaying(goodbye)
		default:
			return
		}
	}
}
