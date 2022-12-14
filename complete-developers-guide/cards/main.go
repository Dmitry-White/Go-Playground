package main

import (
	"fmt"
)

// Error: expected declaration, found name
// name := "Dmitry"

var name string = "Dmitry"

func newCard() string {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Hearts"

	fmt.Println("Author: ", name)

	return card
}

func main() {
	cards := newDeck()

	cards.print()
	fmt.Println("----------------")

	hand, remainingDeck := cards.deal(3)
	hand.print()
	fmt.Println("----------------")
	remainingDeck.print()
	fmt.Println("----------------")

	hand.saveToFile("hand.txt")
	newCards := newDeckFromFile("hand.txt")
	newCards.print()
	fmt.Println("----------------")

	newCards.shuffle()
	newCards.print()
	fmt.Println("----------------")
}
