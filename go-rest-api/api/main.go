package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DB: ", db)

	rows, err := db.Query("SELECT id, name,inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	fmt.Printf("Columns: %+v\n", columns)

	for rows.Next() {
		var p Product
		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Printf("Product: %+v\n", p)
	}

}
