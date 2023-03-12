package main

import (
	"fmt"
	"time"
)

func buffered() {
	// Create an unbuffered channel
	ch := make(chan string, 1)

	// start the greeter to provide a greeting
	go greet(ch)

	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")

	// Receive greeting
	greeting := <-ch

	// Sleep and print
	time.Sleep(5 * time.Second)
	fmt.Println("Greeting received!")
	fmt.Println(greeting)
}
