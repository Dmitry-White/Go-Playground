package main

import (
	"fmt"
)

type Letter struct {
	Symbol      string
	Translation string
}

var Dictionary = []Letter{
	{
		Symbol:      "",
		Translation: "",
	},
}

func englishFor(greek string) string {
	fmt.Println("Not Implemented")
	return ""
}
