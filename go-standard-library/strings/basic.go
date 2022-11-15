package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func basic() {
	s := "The quick brown fox jumps over the lazy dog"

	// TODO: Length of string
	fmt.Println(len(s))

	// TODO: Iterate over each character
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println()

	// TODO: Using operators < > == !=
	fmt.Println("dog" < "cat")
	fmt.Println("dog" < "horse")
	fmt.Println("dog" == "Dog")

	// TODO: Comparing two strings
	result1 := strings.Compare("dog", "cat")
	fmt.Println(result1)
	result2 := strings.Compare("dog", "dog")
	fmt.Println(result2)

	// TODO: EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("Dog", "dog"))

	// TODO: ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s) // Deprecated

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// TODO: New Title approach using Cases and Language
	caser := cases.Title(language.BritishEnglish)
	s4 := caser.String(s)

	fmt.Println(s4)
}
