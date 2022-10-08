package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	expectedDeckLength := 16
	actualDeckLength := len(cards)
	if actualDeckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, actualDeckLength)
	}

	expectedFirstCard := "Ace of Spades"
	actualFirstCard := cards[0]
	if actualFirstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, actualFirstCard)
	}

	expectedLastCard := "Four of Clubs"
	actualLastCard := cards[len(cards)-1]
	if actualLastCard != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, actualLastCard)
	}
}
