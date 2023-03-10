package repo

import (
	"go-applied-concurrency/db"
	"go-applied-concurrency/models"
)

// repo holds all the dependencies required for repo operations
type repo struct {
	products *db.ProductDB
	orders   *db.OrderDB
}

// Repo is the interface we expose to outside packages
type Repo interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
}

// New creates a new Order repo with the correct database dependencies
func New() (Repo, error) {
	p, err := db.NewProducts()
	if err != nil {
		return nil, err
	}
	o := repo{
		products: p,
		orders:   db.NewOrders(),
	}
	return &o, nil
}

// GetAllProducts returns all products in the system
func (r *repo) GetAllProducts() []models.Product {
	return r.products.FindAll()
}

// GetProduct returns the given order if one exists
func (r *repo) GetOrder(id string) (models.Order, error) {
	return r.orders.Find(id)
}

// CreateOrder creates a new order for the given item
func (r *repo) CreateOrder(item models.Item) (*models.Order, error) {
	if err := r.validateItem(item); err != nil {
		return nil, err
	}
	order := models.NewOrder(item)
	r.orders.Upsert(order)
	r.processOrders(&order)
	return &order, nil
}
