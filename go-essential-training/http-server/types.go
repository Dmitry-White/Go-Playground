package main

type Routes struct {
	HEALTH string
	CHECK  string
	MATH   string
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}
