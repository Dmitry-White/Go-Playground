package main

import "os"

const FILE_PATH_1 = "datafile.txt"
const FILE_PATH_2 = "testfile.txt"

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func geFileLength(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return file.Size()
}
