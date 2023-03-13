package main

import (
	"fmt"
	"time"
)

func ranged() {
	// Create a bidirectional channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greetRange(ch)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	// Main goroutine will end up in a deadlock
	// since it has no idea when the goroutine has finished
	// and therefore does not no when to exit the loop
	// being stuck on receving channel message line
	for {
		// Receive greetings
		greeting := <-ch

		// Sleep and print
		time.Sleep(time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}
