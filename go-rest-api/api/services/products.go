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
		product := data.Product{}
		err := rows.Scan(&product.ID, &product.ProductCode, &product.Name, &product.Inventory, &product.Price, &product.Status)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	log.Printf("Products: %+v\n", products)

	return &products, err
}

func GetProduct(db *sql.DB, product *data.Product) error {
	row := db.QueryRow(
		"SELECT name, productCode, inventory, price, status FROM products WHERE id =?",
		product.ID,
	)
	err := row.Scan(&product.Name, &product.ProductCode, &product.Inventory, &product.Price, &product.Status)
	if err != nil {
		return err
	}
	log.Printf("Product: %+v\n", product)

	return nil
}

func CreateProduct(db *sql.DB, product *data.Product) error {
	res, err := db.Exec(
		"INSERT INTO products (name, productCode, inventory, price, status) VALUES(?, ?, ?, ?, ?)",
		product.Name,
		product.ProductCode,
		product.Inventory,
		product.Price,
		product.Status,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)

	return nil
}
