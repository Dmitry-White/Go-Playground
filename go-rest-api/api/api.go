package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	db, err := sql.Open("sqlite3", "./api/practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("DB: %+v\n", db)
	app.DB = db
}

func (app *App) Run() {
	router := mux.NewRouter()
	fmt.Printf("Router: %+v\n", router)

	router.HandleFunc("/handle", app.Handle).Methods("GET")
	router.HandleFunc("/", getRequest).Methods("GET")
	router.HandleFunc("/", postRequest).Methods("POST")
	router.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", router)
	fmt.Printf("Server started and listening on localhost%s\n", app.Port)

	err := http.ListenAndServe(app.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
