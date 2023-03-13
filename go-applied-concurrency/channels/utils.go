package main

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello!", "你好！", "Привет!"}
var goodbyes = []string{"Bye!", "再见！", "Пока!"}

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

// Unidirectional send-only channels will be closed upon completion
func greetRangeClosed(ch chan<- string) {
	fmt.Println("Greeter ready!")

	for _, greeting := range greetings {
		ch <- greeting
	}
	close(ch)

	fmt.Println("Greeter completed!")
}

func greetSelected(sayings []string, ch chan<- string) {
	fmt.Println("Greeter ready!")

	for _, saying := range sayings {
		ch <- saying
	}
	close(ch)

	fmt.Println("Greeter completed!")
}

func printSaying(saying string) {
	time.Sleep(time.Second)
	fmt.Println("Greeting received!", saying)
}
