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

func exercisePattern(name string) (AAA, error) {
	fmt.Println("Exercising pattern...")

	time.Sleep(time.Second)

	if name != "AAA" {
		err := fmt.Errorf("error: wrong pattern %s", name)
		return AAA{}, err
	}

	res := AAA{
		A_1: "Arrange",
		A_2: "Act",
		A_3: "Assert",
	}

	fmt.Println("Done")
	return res, nil
}

func main() {
	res, err := exercisePattern("AAAA")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v", res)
}
