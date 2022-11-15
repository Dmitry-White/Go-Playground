package main

import (
	"bufio"
	"fmt"
	"os"
)

func stdin() {
	// TODO: Using bufio and os
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your input: ")
	s, _ := reader.ReadString('\n')
	fmt.Print("Result:", s)
}
