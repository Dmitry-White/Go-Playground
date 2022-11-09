package main

import "fmt"

func greeting(name, message string) (salutation string) {
	// The variable name specified in the return type is implicitely declared
	salutation = fmt.Sprintf("%s, %s.", message, name)
	return salutation
}

func main() {
	message := greeting("General Kenobi", "Hello there")
	fmt.Println(message)
}
