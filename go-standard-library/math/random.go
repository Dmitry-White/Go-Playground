package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() {
	// TODO: Initialize the random seed to an unknown value
	rand.Seed(time.Now().UnixNano())

	// TODO: Generate random Integer numbers
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))

	// TODO: Generate random Floating Point numbers
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	// TODO: Use the Perm function to create permutations
	s := []string{"apple", "pear", "grape", "orange", "kiwi", "melon"}
	indexes := rand.Perm(len(s))
	fmt.Println(indexes)

	for i := 0; i < len(indexes); i++ {
		fmt.Print(s[indexes[i]], " ")
	}
	fmt.Println()

	// TODO: Generate a random integer between a and b
	const a, b = 10, 50

	for i := 0; i < 10; i++ {
		n := a + rand.Intn(b-a+1)
		fmt.Print(n, " ")
	}
	fmt.Println()

	// TODO: Generate a random uppercase character
	for i := 0; i < 10; i++ {
		c := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Print(c, " ")
	}
	fmt.Println()
}
