package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// "(d deck)" is a receiver.
// It denotes that any variable anywhere in the code
// of type "deck"
// now has access to "print" via ".dot" notation.
// Work kind of similar to "this" in JavaScript
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]

	return hand, remainingDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	data := []byte(d.toString())
	return ioutil.WriteFile(filename, data, 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	stringifiedData := string(data)

	parsedData := strings.Split(stringifiedData, ",")

	cards := deck(parsedData)

	return cards
}

func (d deck) shuffle() {
	for index := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
