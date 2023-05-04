package main

import (
	"fmt"
	"time"
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func postMetric(m *Metric) (*Metric, error) {
	fmt.Println(m)
	return nil, nil
}

func sendMetrics() interface{} {
	m := Metric{
		Time:   time.Now(),
		CPU:    0.23,
		Memory: 87.32,
	}

	reply, err := postMetric(&m)
	if err != nil {
		return err
	}

	return reply
}
