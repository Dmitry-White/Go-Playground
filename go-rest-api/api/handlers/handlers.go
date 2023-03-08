package handlers

import (
	"fmt"
	"go-rest-api/api/data"
	"log"
	"net/http"
)

func Handle(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := app.DB.Query("SELECT id, name,inventory, price FROM products")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer rows.Close()

		columns, _ := rows.Columns()
		log.Printf("Columns: %+v\n", columns)

		for rows.Next() {
			p := data.OldProduct{}
			rows.Scan(p.Id, &p.Name, &p.Inventory, &p.Price)
			log.Printf("Product: %+v\n", p)
		}
	}
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is GET")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is POST")
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is DELETE")
}
