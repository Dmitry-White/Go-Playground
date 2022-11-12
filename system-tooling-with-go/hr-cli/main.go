package main

import "fmt"

func handleError() {
	fmt.Println("[handleError] Not Implemented")
}

func handleCSVFormat() {
	fmt.Println("[handleCSVFormat] Not Implemented")
	handleError()
}

func handleJSONFormat() {
	fmt.Println("[handleJSONFormat] Not Implemented")
	handleError()
}

func handleOutputFormat() {
	fmt.Println("[handleOutputFormat] Not Implemented")
	handleJSONFormat()
	handleCSVFormat()
}

func handleOutputPath() {
	fmt.Println("[handleOutputPath] Not Implemented")
	handleError()
}

func parseFlags() {
	fmt.Println("[parseFlags] Not Implemented")
	handleError()
}

func collectUsers() {
	fmt.Println("[collectUsers] Not Implemented")
	handleError()
}

func main() {
	parseFlags()
	collectUsers()

	handleOutputPath()
	handleOutputFormat()
}
