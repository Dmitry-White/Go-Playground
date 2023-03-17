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

// RequestReversal reverses an order if one exists and completed
func (r *Repo) RequestReversal(id string) (*models.Order, error) {
	order, err := r.Orders.Find(id)
	if err != nil {
		return nil, fmt.Errorf("no order found, try a different one")
	}

	if order.Status != models.OrderStatus_Completed {
		return nil, fmt.Errorf("only completed orders can be reversed, try a different one")
	}

	order.Request()

	select {
	case r.Reversal <- order:
		r.Orders.Upsert(order)
		return &order, nil
	case <-r.Done:
		return nil, fmt.Errorf("orders app is closed, try again later")
	}
}
