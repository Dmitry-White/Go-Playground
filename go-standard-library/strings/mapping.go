package main

import (
	"fmt"
	"strings"
)

const alphabetLength = 26

func shiftRune(r rune, lower int, upper int, shift int) rune {

	value := int(r) + shift

	if value > upper {
		value -= alphabetLength
	} else if value < lower {
		value += alphabetLength
	}
	return rune(value)
}

func mapping() {
	// The map function returns a copy of a string
	// with the characters modified
	// according to the mapping function
	shift := 2
	s := "The quick brown fox jumped over the lazy dog"

	// TODO: Create the mapping function
	transform := func(shift int) func(rune) rune {
		return func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return shiftRune(r, 65, 91, shift)
			case r >= 'a' && r <= 'z':
				return shiftRune(r, 97, 122, shift)
			}

			return r
		}
	}

	// TODO: Encode the message
	encodeTransform := transform(shift)
	encoded := strings.Map(encodeTransform, s)
	fmt.Println(encoded)

	// TODO: Decode the message
	decodeTransform := transform(-shift)
	decoded := strings.Map(decodeTransform, encoded)
	fmt.Println(decoded)
}
