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

func parseTime(ts string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 PM", ts)
}

func getLaggingStations(reader io.Reader, lag time.Duration) ([]string, error) {

	reply := Reply{}

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&reply)
	if err != nil {
		return nil, err
	}

	checkTime, err := parseTime(reply.LastCheckTime)
	if err != nil {
		return nil, err
	}

	lagging := []string{}
	for _, station := range reply.Stations {
		if station.Status != "Active" {
			continue
		}

		lastCheck, err := parseTime(station.LastCheck.Time)
		if err != nil {
			return nil, err
		}

		if checkTime.Sub(lastCheck) > lag {
			lagging = append(lagging, station.Name)
		}
	}

	return lagging, nil
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
