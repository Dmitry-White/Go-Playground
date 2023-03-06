package api

import (
	"database/sql"
	"log"
)

type product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}

func GetProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	log.Printf("Products: %+v\n", products)

	return products, err
}

func (p *product) GetProduct(db *sql.DB) error {
	row := db.QueryRow(
		"SELECT name, productCode, inventory, price, status FROM products WHERE id =?",
		p.ID,
	)
	err := row.Scan(&p.Name, &p.ProductCode, &p.Inventory, &p.Price, &p.Status)
	if err != nil {
		return err
	}
	log.Printf("Product: %+v\n", p)

	return nil
}

func (p *product) CreateProduct(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO products (name, productCode, inventory, price, status) VALUES(?, ?, ?, ?, ?)",
		p.Name,
		p.ProductCode,
		p.Inventory,
		p.Price,
		p.Status,
	)
	if err != nil {
		return err
	}
	log.Printf("Response: %+v\n", res)

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)

	return nil
}
