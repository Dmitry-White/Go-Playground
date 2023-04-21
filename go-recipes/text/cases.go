package main

import (
	"strings"
)

type Letter struct {
	Symbol      string
	Translation string
}

var dictionary = []Letter{
	{"Σ", "Sigma"},
}

func englishFor(greek string) string {
	for _, letter := range dictionary {
		if strings.EqualFold(letter.Symbol, greek) {
			return letter.Translation
		}
	}
	return "Unknown"
}
