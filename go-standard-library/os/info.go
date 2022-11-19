package main

import (
	"fmt"
	"os"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

const VALID_FILE_PATH = "sample.txt"
const INVALID_FILE_PATH = "sample2.txt"

func info() {
	// TODO: Use the Stat function to get file stats
	stats, err := os.Stat(VALID_FILE_PATH)
	if err != nil {
		panic(err)
	}

	// TODO: Check if a file exists
	exists1 := checkFileExists(VALID_FILE_PATH)
	fmt.Println("File1 exists check: ", exists1)

	exists2 := checkFileExists(INVALID_FILE_PATH)
	fmt.Println("File2 exists check: ", exists2)

	// TODO: Get the file's modification time
	fmt.Println("Modification time: ", stats.ModTime())
	fmt.Println("File Mode: ", stats.Mode())
	fmode := stats.Mode()
	if fmode.IsRegular() {
		fmt.Println("This is a regular file")
	}

	// TODO: Get the file size
	fmt.Println("File Size: ", stats.Size())

	// TODO: Is this a directory
	fmt.Println("This is a directory: ", stats.IsDir())

}
