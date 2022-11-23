package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func temporary() {
	tempContent := []byte("This is some content for the temp file")

	// TODO: Get the default temporary directory path
	tempPath := os.TempDir()
	fmt.Println("Default temp path:", tempPath)

	// TODO: Create a temporary file using ioutil.TempFile
	tempFile1, err := ioutil.TempFile(tempPath, "tempfile_*.txt")
	if err != nil {
		panic(err)
	}

	n1, err := tempFile1.Write(tempContent)
	if err != nil {
		panic(err)
	}

	fmt.Println("Temp content 1 size written:", n1)

	// TODO: Create a temporary file using os.CreateTemp
	tempFile2, err := os.CreateTemp(tempPath, "tempfile_*.txt")
	if err != nil {
		panic(err)
	}

	n2, err := tempFile2.Write(tempContent)
	if err != nil {
		panic(err)
	}

	fmt.Println("Temp content 2 size written:", n2)

	// TODO: Read and print the temporary file content
	data1, _ := ioutil.ReadFile(tempFile1.Name())
	data2, _ := ioutil.ReadFile(tempFile2.Name())

	fmt.Println("Temp content 1:", string(data1))
	fmt.Println("Temp content 2:", string(data2))

	// TODO: Create a temporary directory with ioutil.TempDir
	ioTempDir, _ := ioutil.TempDir("", "tempdir*")
	fmt.Println("New IO Temp Dir: ", ioTempDir)

	// TODO: Create a temporary directory with os.TempDir
	osTempDir, _ := ioutil.TempDir("", "tempdir*")
	fmt.Println("New OS Temp Dir: ", osTempDir)

	// TODO: Remove the temporary file when finished
	tempFile1.Close()
	tempFile2.Close()

	tempDirContentBefore, _ := os.ReadDir(tempPath)
	fmt.Printf("\nTemp Dir content length: %d\n", len(tempDirContentBefore))

	for _, dirItem := range tempDirContentBefore {
		fileName := dirItem.Name()

		if strings.Contains(fileName, "tempfile_") {
			filePath := filepath.Join(tempPath, fileName)
			fmt.Println("Filepath:", filePath)

			err := os.Remove(filePath)
			if err != nil {
				panic(err)
			}
		}
	}

	// TODO: Remove temp dir
	os.RemoveAll(ioTempDir)
	os.RemoveAll(osTempDir)

	tempDirContentAfter, _ := os.ReadDir(tempPath)
	fmt.Println("\nTemp Dir content length:", len(tempDirContentAfter))
}
