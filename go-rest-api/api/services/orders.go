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
		order := data.Order{}
		err := rows.Scan(&order.Id, &order.CustomerName, &order.Total, &order.Status)
		if err != nil {
			return nil, err
		}

		orderItems, err := GetOrderItems(db, &order)
		if err != nil {
			return nil, err
		}

		order.Items = *orderItems

		orders = append(orders, order)
	}
	log.Printf("Orders: %+v\n", orders)

	return &orders, err
}

func GetOrder(db *sql.DB, order *data.Order) error {
	row := db.QueryRow(
		"SELECT id, customerName, total, status FROM orders WHERE id =?",
		order.Id,
	)
	err := row.Scan(&order.Id, &order.CustomerName, &order.Total, &order.Status)
	if err != nil {
		return err
	}

	orderItems, err := GetOrderItems(db, order)
	if err != nil {
		return err
	}

	order.Items = *orderItems

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

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	order.Id = int(id)

	return nil
}
