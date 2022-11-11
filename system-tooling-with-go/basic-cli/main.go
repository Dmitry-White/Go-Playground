package main

type Flags struct {
	name     string
	greeting string
	promt    bool
	preview  bool
}

func parseFlags(*Flags) {
	println("Not Implemented")
}

func showUsage(*Flags) {
	println("Not Implemented")
}

func handleDebug(*Flags) {
	println("Not Implemented")
}

func handlePrompt(*Flags) {
	println("Not Implemented")
}

func handleMessage(*Flags) {
	println("Not Implemented")
}

func main() {
	flags := Flags{}

	parseFlags(&flags)

	showUsage(&flags)

	handleDebug(&flags)

	handlePrompt(&flags)

	handleMessage(&flags)
}
