package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatus_New               OrderStatus = "New"
	OrderStatus_Completed         OrderStatus = "Completed"
	OrderStatus_Rejected          OrderStatus = "Rejected"
	OrderStatus_ReversalRequested OrderStatus = "ReversalRequested"
	OrderStatus_Reversed          OrderStatus = "Reversed"
)

const timeFormat = "2006-01-02 15:04:05.000"

type Order struct {
	ID        string      `json:"id,omitempty"`
	Item      Item        `json:"item"`
	Total     *float64    `json:"total,omitempty"`
	Status    OrderStatus `json:"status,omitempty"`
	Error     string      `json:"error,omitempty"`
	CreatedAt string      `json:"createdAt,omitempty"`
}

type Item struct {
	ProductID string `json:"productId"`
	Amount    int    `json:"amount"`
}

func NewOrder(item Item) Order {
	return Order{
		ID:        uuid.New().String(),
		Status:    OrderStatus_New,
		CreatedAt: time.Now().Format(timeFormat),
		Item:      item,
	}
}

func (o *Order) Complete() {
	o.Status = OrderStatus_Completed
}

func (o *Order) Reject() {
	o.Status = OrderStatus_Rejected
}

func (o *Order) Request() {
	o.Status = OrderStatus_ReversalRequested
}

func (o *Order) Reverse() {
	o.Status = OrderStatus_Reversed
}
