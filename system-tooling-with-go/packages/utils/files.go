package utils

import (
	"fmt"
	"log"
	"os"
	"packages/core"
)

func RecordLineQuick(line string) {
	err := os.WriteFile(core.FILE_NAME, []byte(line), 0644)
	if err != nil {
		log.Fatalf("Error: Failed to write to %s, %v", core.FILE_NAME, err)
	}
}

func RecordLine(line string) {
	fmt.Println("Not Implemented")
}

func PullRecords() {
	fmt.Println("Not Implemented")
}
