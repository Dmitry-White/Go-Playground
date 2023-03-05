package api

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("DB: %+v\n", app.DB)
	rows, err := app.DB.Query("SELECT id, name,inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	log.Printf("Columns: %+v\n", columns)

	for rows.Next() {
		var p Product
		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		log.Printf("Product: %+v\n", p)
	}
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is GET")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is POST")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is DELETE")
}
