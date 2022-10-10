package main

import "fmt"

func createLiteralMap() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colours)
}

func createDeclaredMap() {
	// var colours map[string]string
	colours := make(map[string]string)

	// Because maps in Go are statically typed
	// and the key in map should always be of certain type
	// "dot notation" property access is not available
	colours["red"] = "#ff0000"
	colours["green"] = "#4bf745"

	delete(colours, "green")

	fmt.Println(colours)
}

func main() {
	createLiteralMap()
	createDeclaredMap()
}
