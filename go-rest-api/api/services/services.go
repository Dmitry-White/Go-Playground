package services

import (
	"database/sql"
	"go-rest-api/api/data"
	"log"
)

func GetProducts(db *sql.DB) (*[]data.Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []data.Product{}

	for rows.Next() {
		var p data.Product
		err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}
	log.Printf("Products: %+v\n", products)

	return &products, err
}

func GetProduct(db *sql.DB, p *data.Product) error {
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

func CreateProduct(db *sql.DB, p *data.Product) error {
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
