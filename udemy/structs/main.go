package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	zipCode int
}

type employee struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func createAlex() {
	// A struct could be declared and initialized with ordered values
	// alex := person{"Alex", "Anderson"}

	// A struct could be declared and initialized with field names
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// A struct could be declared with zero values
	var alex person

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
	fmt.Println("------------------")

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
	fmt.Println("------------------")
}

func createJim() {
	jim := employee{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v\n", jim)
	fmt.Println("------------------")
}

func main() {
	createAlex()
	createJim()
}
