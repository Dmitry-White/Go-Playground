package main

import "time"

const (
	BASE_URL      = "https://httpbin.org"
	POST_ENDPOINT = "post"
	AUTH_ENDPOINT = "basic-auth"
	SIZE          = 1 << 20
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	Host   string    `json:"host"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}
