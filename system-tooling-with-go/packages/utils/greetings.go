package utils

import "fmt"

func Greeting(name, message string) (salutation string) {
	// The variable name specified in the return type is implicitely declared
	salutation = fmt.Sprintf("%s: %s", name, message)
	return salutation
}
