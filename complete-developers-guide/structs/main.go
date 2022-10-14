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

	// Go passes stuff by value
	// Therefore, when working with mutations,
	// we need to take care of converting values to references (pointers)
	// to enable passing stuff by reference.
	// "value.smth()"" is a syntactic sugar
	// for the actual usage with pointers of "(&value).smth()"
	// Therefore the following 2 lines are equivalent in action
	// (&jim).updateName("Jimmy")
	jim.updateName("Jimmy")

	jim.print()
}
