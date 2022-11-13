package services

import (
	"basic-cli/core"
	"fmt"
	"os"
)

func OutputPreview(message string) {
	fmt.Println(message)
}

func OutputFile(message string) {
	data := []byte(fmt.Sprintf("%s\n", message))

	err := os.WriteFile(core.FILE_NAME, data, 0644)
	if err != nil {
		fmt.Printf("Error: Failed to write: %v", err)
		os.Exit(1)
	}
}
