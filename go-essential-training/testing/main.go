package main

import (
	"fmt"
	"os"
	"time"
)

type AAA struct {
	A_1 string
	A_2 string
	A_3 string
}

var PATTERN = AAA{
	A_1: "Arrange",
	A_2: "Act",
	A_3: "Assert",
}

func exercisePattern(name string) (AAA, error) {
	fmt.Println("Exercising pattern...")

	time.Sleep(time.Second)

	if name != "AAA" {
		err := fmt.Errorf("error: wrong pattern %s", name)
		return AAA{}, err
	}

	fmt.Println("Done")
	return PATTERN, nil
}

func main() {
	res, err := exercisePattern("AAA")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", res)

	verb := getVerb("A_2")

	fmt.Printf("Verb: %v\n", verb)
}
