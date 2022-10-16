package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		fmt.Println("Processing...")
		time.Sleep(11 * time.Second)
		channel <- "Done"
	}()

	select {
	case data := <-channel:
		fmt.Printf("Result: %s", data)
	case currentTime := <-time.After(10 * time.Second):
		fmt.Printf("Timeout: %v", currentTime)
	}
}
