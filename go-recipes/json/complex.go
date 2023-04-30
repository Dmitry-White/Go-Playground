package main

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type Station struct {
	Name      string
	Status    string
	LastCheck struct {
		Time string
	}
}

type Reply struct {
	LastCheckTime string
	Stations      []Station
}

func getLaggingStations(reader io.Reader, lag time.Duration) (interface{}, error) {
	var reply Reply

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func scanStations() interface{} {
	file, err := os.Open("stations.json")
	if err != nil {
		return err
	}
	defer file.Close()

	lagging, err := getLaggingStations(file, time.Minute)
	if err != nil {
		return err
	}

	return lagging
}
