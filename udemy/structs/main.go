package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// A struct could be declared with ordered values
	// alex := person{"Alex", "Anderson"}

	// A struct could be declared with field names
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}
