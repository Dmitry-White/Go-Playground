package main

import (
	"flag"
	"fmt"
	"os"
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

func handleDebug(*Flags) {
	fmt.Println("Not Implemented")
}

func handlePrompt(f *Flags) {
	fmt.Println("Not Implemented")
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

	handleDebug(&flags)

	handlePrompt(&flags)

	handleMessage(&flags)
}
