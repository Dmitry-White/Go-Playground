package services

import (
	"database/sql"
	"go-rest-api/api/data"
	"log"
)

func GetItems(db *sql.DB) (*[]data.OrderItem, error) {
	rows, err := db.Query("SELECT * FROM order_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := []data.OrderItem{}

	for rows.Next() {
		orderItem := data.OrderItem{}
		err := rows.Scan(&orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity)
		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}
	log.Printf("OrderItems: %+v\n", orderItems)

	return &orderItems, err
}

func GetOrderItems(db *sql.DB, order *data.Order) (*[]data.OrderItem, error) {
	rows, err := db.Query(
		"SELECT order_id, product_id, quantity FROM order_items WHERE order_id =?",
		order.Id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := []data.OrderItem{}

	for rows.Next() {
		orderItem := data.OrderItem{}
		err := rows.Scan(&orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity)
		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}
	log.Printf("OrderItems: %+v\n", orderItems)

	return &orderItems, err
}

func CreateOrderItem(db *sql.DB, orderItem *data.OrderItem) error {
	res, err := db.Exec(
		"INSERT INTO order_items (order_id, product_id, quantity) VALUES(?, ?, ?)",
		orderItem.OrderId,
		orderItem.ProductId,
		orderItem.Quantity,
	)
	if err != nil {
		return err
	}
	log.Printf("Response: %+v\n", res)

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	log.Printf("Id: %+v\n", id)

	return nil
}
