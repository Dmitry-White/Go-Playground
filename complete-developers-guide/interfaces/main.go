package main

import "fmt"

// Bot Interface - an Interface Type
// Interface Type ~= Abstract Class
// i.e. can't be instantiated.
// Interface Type != Generic Type
// i.e. it doesn't imply that it will "accept" different types
// A link from Concrete Type to Interface Type is implicit
// i.e. we don't need to explicitly specify that
// some Interface Type works with some Concrete Types or vice versa
type bot interface {
	getGreeting() string
}

// English/Spanish Bot Struct - a Concrete Type
// i.e. can be instantiated
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
