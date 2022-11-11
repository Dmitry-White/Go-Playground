package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	name     string
	greeting string
	prompt   bool
	preview  bool
}

const FILE_NAME = "conversation.txt"

func parseFlags(f *Flags) {
	flag.StringVar(&f.name, "name", "123", "name to use within the message")
	flag.StringVar(&f.greeting, "greeting", "", "phrase to use within the message")
	flag.BoolVar(&f.prompt, "prompt", false, "use prompt to input name and greeting")
	flag.BoolVar(&f.preview, "preview", false, "use preview to output message without writing to a file")

	flag.Parse()
}

func showUsage(f *Flags) {
	// Note:
	// f.prompt is a syntactic sugar for
	// (*f).name
	// to reduce typing
	if !f.prompt && (f.name == "" || f.greeting == "") {
		flag.Usage()
		os.Exit(1)
	}
}

func handleDebug(f *Flags) {
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name: ", f.name)
		fmt.Println("Greeting: ", f.greeting)
		fmt.Println("Prompt: ", f.prompt)
		fmt.Println("Preview: ", f.preview)
	}
}

func handlePrompt(f *Flags) {
	if f.prompt {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Your Name: ")
		inputName, _ := reader.ReadString('\n')
		f.name = strings.TrimSpace(inputName)

		fmt.Print("Your Greeting: ")
		inputGreeting, _ := reader.ReadString('\n')
		f.greeting = strings.TrimSpace(inputGreeting)
	}
}

func handlePreview(message string) {
	fmt.Println(message)
}

func handleWrite(message string) {
	data := []byte(fmt.Sprintf("%s\n", message))

	err := os.WriteFile(FILE_NAME, data, 0644)
	if err != nil {
		fmt.Printf("Error: Failed to write: %v", err)
		os.Exit(1)
	}
}

func handleMessage(f *Flags) {
	message := fmt.Sprintf("%s, %s", f.greeting, f.name)

	if f.preview {
		handlePreview(message)
	} else {
		handleWrite(message)
	}
}

func main() {
	flags := Flags{}

	parseFlags(&flags)

	showUsage(&flags)

	handlePrompt(&flags)

	handleDebug(&flags)

	handleMessage(&flags)
}
