package repo

import "go-applied-concurrency/api/models"

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
	r.Orders.Upsert(order)

	// ProcessOrders(r, &order)
	r.Incomming <- order
	return &order, nil
}
