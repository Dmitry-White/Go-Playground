package main

import "fmt"

func greet(ch chan string) {
	fmt.Println("Greeter ready!")
	ch <- "Hello, World!"
	fmt.Println("Greeter completed!")
}
