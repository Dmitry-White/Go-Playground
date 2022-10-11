package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Since the receiver value is not used,
// it could be omitted
func (englishBot) getGreeting() string {
	// VERY custom logic for getting English greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for getting Spanish greeting
	return "Hola!"
}

// Go does not support function overloading
// therefor we can't use the same function name
// albeit with diffferet parameters

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
