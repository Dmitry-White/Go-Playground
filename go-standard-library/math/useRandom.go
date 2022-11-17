package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func useRandom() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// TODO: Shuffle a sequence of values
	const numString = "one two three four five six"
	words := strings.Fields(numString)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(strings.Join(words, " "))

	// TODO: Generate a random string
	const upLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowLetters = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const specChars = "~=+%^*()[]{}!@#$?|"
	const allChars = upLetters + lowLetters + digits + specChars

	var sb strings.Builder
	legth := 10

	for i := 0; i < legth; i++ {
		charIndex := rand.Intn(len(allChars))
		sb.WriteRune(rune(allChars[charIndex]))
	}
	fmt.Println("String result: ", sb.String())

	// TODO: Generate random string with rules
	buf := make([]byte, legth)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specChars[rand.Intn(len(specChars))]
	buf[2] = upLetters[rand.Intn(len(upLetters))]
	buf[3] = lowLetters[rand.Intn(len(lowLetters))]

	for i := 4; i < legth; i++ {
		buf[i] = allChars[rand.Intn(len(allChars))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	fmt.Println("Rules result: ", string(buf))
}
