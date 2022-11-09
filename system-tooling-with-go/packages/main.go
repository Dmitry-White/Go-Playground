package main

import (
	"bufio"
	"fmt"
	"os"
	"packages/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 1st person: ")
	firstPersonSays := utils.ProcessEntries(reader)

	fmt.Print("Enter 2nd person: ")
	secondPersonSays := utils.ProcessEntries(reader)

	fmt.Println("Your conversation:")
	fmt.Println(firstPersonSays)
	fmt.Println(secondPersonSays)
}
