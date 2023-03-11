package repo

import "go-applied-concurrency/models"

// GetAllProducts returns all products in the system
func (r *Repo) GetAllProducts() []models.Product {
	return r.Products.FindAll()
}
