package data

type Order struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customerName"`
	Total        int    `json:"total"`
	Status       string `json:"status"`
}
