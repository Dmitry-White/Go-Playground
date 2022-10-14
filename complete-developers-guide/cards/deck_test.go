package main

import (
	"os"
	"testing"
)

const TEST_FILE_NAME = "_decktesting"
const TEST_CARD_LENGTH = 16

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	expectedDeckLength := TEST_CARD_LENGTH
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

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(TEST_FILE_NAME)

	cards := newDeck()

	cards.saveToFile(TEST_FILE_NAME)
	newCards := newDeckFromFile(TEST_FILE_NAME)

	expectedDeckLength := TEST_CARD_LENGTH
	actualDeckLength := len(newCards)
	if actualDeckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, actualDeckLength)
	}

	os.Remove(TEST_FILE_NAME)
}
