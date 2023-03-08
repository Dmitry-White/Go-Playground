package tests

import (
	"go-rest-api/api"
	"go-rest-api/api/data"
)

var app api.App = api.App{
	App: &data.App{
		Port:           ":8000",
		DataSourceName: "../practiceit.db",
	},
}

const productTableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id INT NOT NULL PRIMARY KEY AUTOINCREMENT,
	productCode VARCHAR(25) NOT NULL,
	name VARCHAR(25) NOT NULL,
	inventory INT NOT NULL,
	price INT NOT NULL,
	status VARCHAR(64) NOT NULL
)`
const productTableClearingQuery = "DELETE FROM products"
const productTableDeletionQuery = "DELETE FROM sqlite_sequence WHERE name = 'products'"
const productCreationQuery = `INSERT INTO products 
(
	name,
	productCode,
	inventory,
	price,
	status
) VALUES(
	"ProductTest",
	"TEST12345",
	1,
	1,
	"testing"
)`

const orderTableCreationQuery = `CREATE TABLE IF NOT EXISTS orders
(
	id INT NOT NULL PRIMARY KEY AUTOINCREMENT,
	customerName VARCHAR(256) NOT NULL,
	total INT NOT NULL,
	status VARCHAR(64) NOT NULL
)`
const orderTableClearingQuery = "DELETE FROM orders"
const orderTableDeletionQuery = "DELETE FROM sqlite_sequence WHERE name = 'orders'"
const orderCreationQuery = `INSERT INTO orders 
(
	customerName,
	total,
	status
) VALUES(
	"Mikki Mouse",
	200,
	"testing"
)`

const orderItemTableCreationQuery = `CREATE TABLE IF NOT EXISTS order_items
(
	order_id INT,
	product_id INT,
	quantity INT NOT NULL,
	FOREIGN KEY (order_id) REFERENCES orders (id)
	FOREIGN KEY (product_id) REFERENCES products (id)
	PRIMARY KEY (order_id, product_id)
)`
const orderItemTableClearingQuery = "DELETE FROM order_items"
const orderItemTableDeletionQuery = "DELETE FROM sqlite_sequence WHERE name = 'order_items'"
const orderItemCreationQuery = `INSERT INTO order_items 
(
	order_id,
	product_id,
	quantity
) VALUES(
	1,
	1,
	1
)`
