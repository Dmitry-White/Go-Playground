package main

import "fmt"

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
