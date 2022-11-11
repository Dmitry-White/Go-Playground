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
	promt    bool
	preview  bool
}

func parseFlags(f *Flags) {
	flag.StringVar(&f.name, "name", "123", "name to use within the message")
	flag.StringVar(&f.greeting, "greeting", "", "phrase to use within the message")
	flag.BoolVar(&f.promt, "promt", false, "use prompt to input name and greeting")
	flag.BoolVar(&f.preview, "preview", false, "use preview to output message without writing to conversation.txt")

	flag.Parse()
}

func showUsage(f *Flags) {
	// Note:
	// f.promt is a syntactic sugar for
	// (*f).name
	// to reduce typing
	if !f.promt && (f.name == "" || f.greeting == "") {
		flag.Usage()
		os.Exit(1)
	}
}

func handleDebug(f *Flags) {
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name: ", f.name)
		fmt.Println("Greeting: ", f.greeting)
		fmt.Println("Prompt: ", f.promt)
		fmt.Println("Preview: ", f.preview)
	}
}

func handlePrompt(f *Flags) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Name: ")
	inputName, _ := reader.ReadString('\n')
	f.name = strings.TrimSpace(inputName)

	fmt.Print("Your Greeting: ")
	inputGreeting, _ := reader.ReadString('\n')
	f.greeting = strings.TrimSpace(inputGreeting)
}

func handlePreview(message string) {
	fmt.Println(message)
}

func handleWrite(message string) {
	fmt.Println("Not Implemented")
}

func handleMessage(f *Flags) {
	message := ""

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
