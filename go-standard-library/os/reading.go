package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func reading() {
	// TODO: Use the ReadFile function to read an entire file
	content, err := ioutil.ReadFile(FILE_PATH_1)
	if err != nil {
		panic(err)
	}

	// TODO: ReadFile reads the data as bytes, we have to convert to a string
	fmt.Println(string(content))

	// TODO: The Read function can perform a generic read
	const BUFF_SIZE = 5
	file, err := os.Open(FILE_PATH_2)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff := make([]byte, BUFF_SIZE)

	for {
		n, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break

		}

		fmt.Println("Bytes read:", n)
		fmt.Println("Content:", string(buff))
	}

	// TODO: Get the length of a file
	fileLength := geFileLength(FILE_PATH_2)
	fmt.Println("File length: ", fileLength)

	// TODO: Use ReadAt to read from a specific index
	buffSize := 10
	buff2 := make([]byte, buffSize)
	readRange := fileLength - int64(buffSize)
	n, err := file.ReadAt(buff2, readRange)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data: %s, Length: %d\n", string(buff2), n)
}
