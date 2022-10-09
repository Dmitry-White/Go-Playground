package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// A struct could be declared and initialized with ordered values
	// alex := person{"Alex", "Anderson"}

	// A struct could be declared and initialized with field names
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// A struct could be declared with zero values
	var alex person

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
