package data

type OrderItem struct {
	OrderId   int `json:"orderId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
