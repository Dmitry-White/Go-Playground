package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (app *App) AllProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("AllProducts Request: %+v\n", r)

	products, err := GetProducts(app.DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Fprintln(w, products)
}

func (app *App) FetchProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("FetchProduct Request: %+v\n", r)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	p := product{
		ID: id,
	}
	err := p.GetProduct(app.DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Fprintln(w, p)
}
