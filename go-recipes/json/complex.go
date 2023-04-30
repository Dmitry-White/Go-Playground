package main

import (
	"fmt"
	"os"
)

func scanStations() error {
	file, err := os.Open("stations.json")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("File: ", file)

	return nil
}
