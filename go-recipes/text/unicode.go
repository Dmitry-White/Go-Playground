package main

import (
	"unicode/utf8"
)

func sentenceLength(words []string) int {
	count := 0
	for _, word := range words {
		count += utf8.RuneCountInString(word)
	}

	count += len(words) - 1
	return count
}
