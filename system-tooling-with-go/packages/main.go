package main

import (
	"bufio"
	"fmt"
	"os"
	"packages/core"
	"packages/utils"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 1st person: ")
	firstPersonEntry, _ := reader.ReadString('\n')
	trimmedFirstPersonEntry := strings.TrimSpace(firstPersonEntry)
	firstPersonName := core.Names[trimmedFirstPersonEntry]
	firstPersonMessage := core.Messages[trimmedFirstPersonEntry]
	firstPersonSays := utils.Greeting(firstPersonName, firstPersonMessage)

	fmt.Print("Enter 2nd person: ")
	secondPersonEntry, _ := reader.ReadString('\n')
	trimmedSecondPersonEntry := strings.TrimSpace(secondPersonEntry)
	secondPersonName := core.Names[trimmedSecondPersonEntry]
	secondPersonMessage := core.Messages[trimmedSecondPersonEntry]
	secondPersonSays := utils.Greeting(secondPersonName, secondPersonMessage)

	fmt.Println("Your conversation:")
	fmt.Println(firstPersonSays)
	fmt.Println(secondPersonSays)
}
