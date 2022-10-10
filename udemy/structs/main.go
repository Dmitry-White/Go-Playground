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
	// contactInfo contactInfo
	// OR simply
	contactInfo
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

func (e employee) print() {
	fmt.Printf("%+v\n", e)
	fmt.Println("------------------")
}

func (employeePointer *employee) updateName(newFirstName string) {
	(*employeePointer).firstName = newFirstName
}

func createJim() employee {
	jim := employee{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	return jim
}

func main() {
	createAlex()
	jim := createJim()
	jim.print()

	(&jim).updateName("Jimmy")

	jim.print()
}
