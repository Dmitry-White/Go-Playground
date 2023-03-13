package repo

import (
	"go-applied-concurrency/api/db"
	"go-applied-concurrency/api/models"
	"sync"
)

// Repo holds all the dependencies required for Repo operations
type Repo struct {
	Products  *db.ProductDB
	Orders    *db.OrderDB
	Mutex     sync.Mutex
	Incomming chan models.Order
}

// IRepo is the interface we expose to outside packages
type IRepo interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
}

// New creates a new Order&Products Repo with the correct database dependencies
func New(dbPath string) (IRepo, error) {
	p, err := db.NewProducts(dbPath)
	if err != nil {
		return nil, err
	}
	o := Repo{
		Products:  p,
		Orders:    db.NewOrders(),
		Incomming: make(chan models.Order),
	}

	go ProcessOrders(&o)

	return &o, nil
}
