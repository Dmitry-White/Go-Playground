package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MetricResponse struct {
	Json Metric
}

const (
	TIMEOUT = 5 * time.Second
)

func postMetric(m Metric) (*MetricResponse, error) {
	fmt.Println("Metric:", m)

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	fmt.Println("Bytes:", data)

	url := fmt.Sprintf("%s/%s", BASE_URL, POST_ENDPOINT)
	body := bytes.NewReader(data)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	fmt.Println("Request:", request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d %s", response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response)

	result := MetricResponse{}
	reader := io.LimitReader(response.Body, SIZE)
	decoder := json.NewDecoder(reader)

	decoderErr := decoder.Decode(&result)
	if err != nil {
		return nil, decoderErr
	}
	fmt.Printf("Result: %#v\n", result)

	return &result, nil
}

func sendMetrics() interface{} {
	m := Metric{
		Time:   time.Now(),
		CPU:    0.23,
		Memory: 87.32,
	}

	reply, err := postMetric(m)
	if err != nil {
		return err
	}

	return reply
}
