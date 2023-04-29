package main

import (
	"encoding/json"
	"os"
)

func parseTemperature(file *os.File) (interface{}, error) {
	var record interface{}

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func getTemperature() interface{} {
	file, err := os.Open("record.json")
	if err != nil {
		return err
	}
	defer file.Close()

	record, err := parseTemperature(file)
	if err != nil {
		return err
	}

	return record
}
