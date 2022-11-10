package utils

import (
	"fmt"
	"log"
	"os"
	"packages/core"
)

func RecordLineQuick(line string) {
	data := []byte(fmt.Sprintf("%s\n", line))

	err := os.WriteFile(core.FILE_NAME, data, 0644)
	if err != nil {
		log.Fatalf("Error: Failed to write to %s, %v", core.FILE_NAME, err)
	}
}

func RecordLine(line string) {
	data := []byte(fmt.Sprintf("%s\n", line))

	file, err := os.OpenFile(core.FILE_NAME, os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error: Failed to open %s, %v", core.FILE_NAME, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Error: Failed to write to %s, %v", core.FILE_NAME, err)
	}
}

func PullRecords() {
	fmt.Println("Not Implemented")
}
