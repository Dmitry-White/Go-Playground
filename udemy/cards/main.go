package main

import "fmt"

// Error: expected declaration, found name
// name := "Dmitry"

var name string = "Dmitry"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Hearts"

	fmt.Println(card)
}
