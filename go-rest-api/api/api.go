package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

type App struct {
	DB   *sql.DB
	Port string
}

func (app *App) Init() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("DB: ", db)
	app.DB = db
}

func (app *App) Run() {
	http.HandleFunc("/", app.Handle)
	fmt.Printf("Server started and listening on localhost%s\n", app.Port)

	err := http.ListenAndServe(app.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Handle(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, name,inventory, price FROM products")
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
