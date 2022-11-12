package main

import (
	"errors"
	"fmt"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Home  string `json:"home"`
	Shell string `json:"shell"`
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func handleCSVFormat() {
	fmt.Println("[handleCSVFormat] Not Implemented")
	err := errors.New("sample error")
	handleError(err)
}

func handleJSONFormat() {
	fmt.Println("[handleJSONFormat] Not Implemented")
	handleError(nil)
}

func handleOutputFormat() {
	fmt.Println("[handleOutputFormat] Not Implemented")
	handleJSONFormat()
	handleCSVFormat()
}

func handleOutputPath() {
	fmt.Println("[handleOutputPath] Not Implemented")
	handleError(nil)
}

func parseFlags() {
	fmt.Println("[parseFlags] Not Implemented")
	handleError(nil)
}

func collectUsers() {
	fmt.Println("[collectUsers] Not Implemented")
	handleError(nil)
}

func main() {
	parseFlags()
	collectUsers()

	handleOutputPath()
	handleOutputFormat()
}
