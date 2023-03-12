package repo

import "go-applied-concurrency/api/models"

// GetAllProducts returns all products in the system
func (r *Repo) GetAllProducts() []models.Product {
	return r.Products.FindAll()
}
