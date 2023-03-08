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

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id INT NOT NULL PRIMARY KEY AUTOINCREMENT,
	productCode VARCHAR(25) NOT NULL,
	name VARCHAR(25) NOT NULL,
	inventory INT NOT NULL,
	price INT NOT NULL,
	status VARCHAR(64) NOT NULL
)`
const tableClearingQuery = "DELETE FROM products"
const tableDeletionQuery = "DELETE FROM sqlite_sequence WHERE name = 'products'"

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
