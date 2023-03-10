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

// New creates a new Order&Products repo with the correct database dependencies
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
