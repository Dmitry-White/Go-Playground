package main

import "fmt"

type deck []string

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
