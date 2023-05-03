package main

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

func handleRequest(data []byte) (interface{}, error) {
	return data, nil
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
