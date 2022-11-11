package main

import "fmt"

type Flags struct {
	name     string
	greeting string
	promt    bool
	preview  bool
}

func parseFlags(*Flags) {
	fmt.Println("Not Implemented")
}

func showUsage(*Flags) {
	fmt.Println("Not Implemented")
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
