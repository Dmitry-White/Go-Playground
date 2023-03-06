package services

import (
	"database/sql"
	"go-rest-api/api/data"
	"log"
)

func GetOrders(db *sql.DB) (*[]data.Order, error) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []data.Order{}

	for rows.Next() {
		var order data.Order
		err := rows.Scan(&order.Id, &order.CustomerName, &order.Total, &order.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	log.Printf("Orders: %+v\n", orders)

	return &orders, err
}

func GetOrder(db *sql.DB, order *data.Order) error {
	row := db.QueryRow(
		"SELECT customerName, total, status FROM orders WHERE id =?",
		order.Id,
	)
	err := row.Scan(&order.CustomerName, &order.Total, &order.Status)
	if err != nil {
		return err
	}
	log.Printf("Order: %+v\n", order)

	return nil
}

func CreateOrder(db *sql.DB, order *data.Order) error {
	res, err := db.Exec(
		"INSERT INTO orders (customerName, total, status) VALUES(?, ?, ?)",
		order.CustomerName,
		order.Total,
		order.Status,
	)
	if err != nil {
		return err
	}
	log.Printf("Response: %+v\n", res)

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	order.Id = int(id)

	return nil
}
