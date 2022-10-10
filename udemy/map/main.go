package main

import "fmt"

func createLiteralMap() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colours)
}

func main() {
	createLiteralMap()
}
