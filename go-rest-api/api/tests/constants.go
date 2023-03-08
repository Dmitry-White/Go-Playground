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

const tableProductCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id INT NOT NULL PRIMARY KEY AUTOINCREMENT,
	productCode VARCHAR(25) NOT NULL,
	name VARCHAR(25) NOT NULL,
	inventory INT NOT NULL,
	price INT NOT NULL,
	status VARCHAR(64) NOT NULL
)`
const tableProductClearingQuery = "DELETE FROM products"
const tableProductDeletionQuery = "DELETE FROM sqlite_sequence WHERE name = 'products'"
