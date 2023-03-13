package main

import "fmt"

var greetings = []string{"Hello!", "你好！", "Привет!"}

func greet(ch chan string) {
	fmt.Println("Greeter ready!")
	ch <- "Hello, World!"
	fmt.Println("Greeter completed!")
}

// Bidirectional channels will be implicitly cast to unidirectional
func greetUni(ch chan<- string) {
	fmt.Println("Greeter ready!")
	ch <- "Hello, World!"
	fmt.Println("Greeter completed!")
}

// Bidirectional channels will be implicitly cast to unidirectional
func greetRange(ch chan<- string) {
	fmt.Println("Greeter ready!")

	for _, greeting := range greetings {
		ch <- greeting
	}
	fmt.Println("Greeter completed!")
}
