package main

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
	cards := newDeck()

	cards.print()

	hand, remainingDeck := cards.deal(3)

	hand.print()
	remainingDeck.print()
}
