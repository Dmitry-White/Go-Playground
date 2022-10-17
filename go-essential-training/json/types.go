package main

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Response struct {
	Ok      bool    `json:"ok"`
	Balance float64 `json:"balance"`
}
