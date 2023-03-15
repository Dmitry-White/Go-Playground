package repo

import (
	"fmt"
	"go-applied-concurrency/api/models"
)

// GetProduct returns the given order if one exists
func (r *Repo) GetOrder(id string) (models.Order, error) {
	return r.Orders.Find(id)
}

// CreateOrder creates a new order for the given item
func (r *Repo) CreateOrder(item models.Item) (*models.Order, error) {
	if err := r.validateItem(item); err != nil {
		return nil, err
	}
	order := models.NewOrder(item)

	// ProcessOrders(r, &order)
	// r.Incoming <- order

	select {
	case r.Incoming <- order:
		r.Orders.Upsert(order)
		return &order, nil
	case <-r.Done:
		return nil, fmt.Errorf("orders app is closed, try again later")
	}
}
