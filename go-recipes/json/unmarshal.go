package main

import (
	"encoding/json"
	"os"
	"time"
)

type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"` // celsius
	Rain    float64 // millimeter
}

func parseTemperature(file *os.File) (Record, error) {
	var record Record

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&record)
	if err != nil {
		return Record{}, err
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
