package tests

const tableProductCreationQuery = `CREATE TABLE IF NOT EXISTS product
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
