package utils

import (
	"bufio"
	"cobraed-cli/core"
	"cobraed-cli/services"
	"fmt"
	"os"
	"strings"
)

func HandlePrompt(f *core.Flags) {
	if f.Prompt {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Your Name: ")
		inputName, _ := reader.ReadString('\n')
		f.Name = strings.TrimSpace(inputName)

		fmt.Print("Your Greeting: ")
		inputGreeting, _ := reader.ReadString('\n')
		f.Greeting = strings.TrimSpace(inputGreeting)
	}
}

func HandleDebug(f *core.Flags) {
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name: ", f.Name)
		fmt.Println("Greeting: ", f.Greeting)
		fmt.Println("Prompt: ", f.Prompt)
		fmt.Println("Preview: ", f.Preview)
	}
}

func HandleMessage(f *core.Flags) {
	message := fmt.Sprintf("%s, %s", f.Greeting, f.Name)

	if f.Preview {
		services.OutputPreview(message)
	} else {
		services.OutputFile(message)
	}
}
