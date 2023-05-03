package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// StartJob is a request to start a job
type StartJob struct {
	Type  string
	User  string
	Count int
}

// JobStatus is a request for job status
type JobStatus struct {
	Type string
	ID   string
}

func handleStartJob(raw map[string]interface{}) StartJob {
	startJob := StartJob{}

	err := mapstructure.Decode(raw, &startJob)
	if err != nil {
		fmt.Println("Can't decode raw data", raw)
	}

	return startJob
}

func handleJobStatus(raw map[string]interface{}) JobStatus {
	jobStatus := JobStatus{}

	err := mapstructure.Decode(raw, &jobStatus)
	if err != nil {
		fmt.Println("Can't decode raw data", raw)
	}

	return jobStatus
}

func handleRequest(data []byte) (interface{}, error) {
	temp := map[string]interface{}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return nil, err
	}

	field, ok := temp["type"]
	if !ok {
		return nil, fmt.Errorf("no type field")
	}

	typeField, ok := field.(string)
	if !ok {
		return nil, fmt.Errorf("type is not a string")
	}

	switch typeField {
	case "start":
		return handleStartJob(temp), nil
	case "status":
		return handleJobStatus(temp), nil
	}

	return nil, fmt.Errorf("unknown request type")
}

func handleRequests() interface{} {
	firstRequest := []byte(`{"type": "start", "user": "joe", "count": 7}`)
	firstResponse, err := handleRequest(firstRequest)
	if err != nil {
		return err
	}

	secondRequest := []byte(`{"type": "status", "id": "seven"}`)
	secondResponse, err := handleRequest(secondRequest)
	if err != nil {
		return err
	}

	return []interface{}{firstResponse, secondResponse}
}
