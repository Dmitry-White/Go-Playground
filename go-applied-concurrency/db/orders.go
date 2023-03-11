package db

import (
	"fmt"
	"sync"

	"go-applied-concurrency/models"
)

type OrderDB struct {
	// Naive approach using built-in Map
	// placedOrders map[string]models.Order
	placedOrders sync.Map
}

// NewOrders creates a new empty order service
func NewOrders() *OrderDB {
	return &OrderDB{
		// Naive approach using built-in Map
		// placedOrders: make(map[string]models.Order),
	}
}

// Find order for a given id, if one exists
func (o *OrderDB) Find(id string) (models.Order, error) {
	// Naive approach using built-in Map
	// order, ok := o.placedOrders[id]
	order, ok := o.placedOrders.Load(id)
	if !ok {
		return models.Order{}, fmt.Errorf("no order found for %s order id", id)
	}

	return toOrder(order), nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	// Naive approach using built-in Map
	// o.placedOrders[order.ID] = order
	o.placedOrders.Store(order.ID, order)
}
