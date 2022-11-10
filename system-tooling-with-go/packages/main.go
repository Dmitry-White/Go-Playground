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
	utils.RecordLineQuick(firstPersonSays)

	fmt.Print("Enter 2nd person: ")
	secondPersonSays := utils.ProcessEntries(reader)
	utils.RecordLine(secondPersonSays)

	fmt.Println("Your conversation:")
	utils.PullRecords()
}
