package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newCard()
	deckLength := len(cards)

	if deckLength != 13 {
		t.Errorf("Expected deck length of 13, but got %v", deckLength)
	}
}
