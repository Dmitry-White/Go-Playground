package main

import (
	"fmt"
	"os"
)

func directory() {
	// TODO: Create a directory
	os.Mkdir("utils", os.ModePerm)
	os.Mkdir("temp", os.ModePerm)

	// TODO: Create a deep directory
	os.MkdirAll("misc/temp", os.ModePerm)

	// TODO: Remove will remove an item
	os.Remove("utils")

	// TODO: RemoveAll will remove an item and all children
	os.RemoveAll("misc")

	// TODO: Get the current working directory
	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)

	// TODO: Get the directory of the current process
	exedir, _ := os.Executable()
	fmt.Println("Path to exe:", exedir)

	// TODO: Read the contents of a directory
	constents, _ := os.ReadDir(".")
	for _, file := range constents {
		fmt.Println(file.Name(), file.IsDir())
	}
}
