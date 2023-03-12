package db

import (
	"fmt"
	"sort"
	"sync"

	"go-applied-concurrency/api/models"
)

type ProductDB struct {
	// Naive approach using built-in Map
	// products map[string]models.Product
	products sync.Map
}

// NewProducts creates a new empty product DB
func NewProducts(dbPath string) (*ProductDB, error) {
	p := &ProductDB{
		// Naive approach using built-in Map
		// products: make(map[string]models.Product),
	}
	// load start position
	if err := ImportProducts(&p.products, dbPath); err != nil {
		return nil, err
	}

	return p, nil
}

// Exists checks whether a product with a give id exists
func (p *ProductDB) Exists(id string) error {
	// Naive approach using built-in Map
	if _, ok := /*p.products[id]*/ p.products.Load(id); !ok {
		return fmt.Errorf("no product found for id %s", id)
	}

	return nil
}

// Find returns a given product if one exists
func (p *ProductDB) Find(id string) (models.Product, error) {
	// Naive approach using built-in Map
	// prod, ok := p.products[id]
	prod, ok := p.products.Load(id)
	if !ok {
		return models.Product{}, fmt.Errorf("no product found for id %s", id)
	}

	return toProduct(prod), nil
}

// Upsert creates or updates a product in the orders DB
func (p *ProductDB) Upsert(prod models.Product) {
	// Naive approach using built-in Map
	// p.products[prod.ID] = prod
	p.products.Store(prod.ID, prod)
}

// FindAll returns all products in the system
func (p *ProductDB) FindAll() []models.Product {
	var allProducts []models.Product
	// Naive approach using built-in Map
	// for _, product := range p.products {
	// 	allProducts = append(allProducts, product)
	// }
	p.products.Range(func(_, product interface{}) bool {
		allProducts = append(allProducts, toProduct(product))
		return true
	})
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].ID < allProducts[j].ID
	})
	return allProducts
}
