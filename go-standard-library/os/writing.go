package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const FILE_PATH_1 = "datafile.txt"
const FILE_PATH_2 = "testfile.txt"

func wirteFileData() {
	// TODO: Create a new file
	file, err := os.Create(FILE_PATH_2)
	if err != nil {
		panic(err)
	}

	// TODO: It is idiomatic in Go to defer the close after you open the file
	defer file.Close()

	// TODO: Get the Name of the file
	fmt.Println("File name: ", file.Name())

	// TODO: Write some string data to the file
	file.WriteString("This is some text\n")

	// TODO: Write som individual bytes to the file
	data := []byte{'a', 'b', 'c', '\n'}
	file.Write(data)

	// TODO: The Sync function forces the data to be written out
	file.Sync()
}

func appendFileData(filname string, data string) {
	// TODO: Use the lower-level OpenFile function to open the file in append mode
	file, err := os.OpenFile(filname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(data)
}

func writing() {
	// TODO: Simple example: dump some data to a file
	data1 := []byte("This is some text data\n")
	ioutil.WriteFile(FILE_PATH_1, data1, 0644)

	// TODO: More complex example: granular writing of data
	if !checkFileExists(FILE_PATH_2) {
		wirteFileData()
	} else {
		os.Truncate(FILE_PATH_2, 10)
		fmt.Printf("%s has been trimmed down!\n", FILE_PATH_2)
	}

	// TODO: Append data to a file
	appendFileData(FILE_PATH_2, "This data was appended")
}
