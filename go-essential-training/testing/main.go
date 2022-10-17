package main

import (
	"fmt"
	"os"
	"time"
)

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
