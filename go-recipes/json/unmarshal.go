package main

import (
	"fmt"
	"os"
)

func getTemperature() error {
	file, err := os.Open("record.json")
	if err != nil {
		return err
	}
	fmt.Println("File", file)
	return nil
}
