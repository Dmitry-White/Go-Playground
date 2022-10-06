package main

import "fmt"

// Error: expected declaration, found name
// name := "Dmitry"

var name string = "Dmitry"

func newCard() string {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Hearts"

	return card
}

func main() {
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Five of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
