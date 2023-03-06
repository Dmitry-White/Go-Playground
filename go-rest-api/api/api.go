package api

import (
	"database/sql"
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
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func (app *App) Init() {
	db, err := sql.Open("sqlite3", "./api/practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("DB: %+v\n", db)

	router := mux.NewRouter()
	log.Printf("Router: %+v\n", router)

	app.DB = db
	app.Router = router
	app.initRoutes()
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/handle", app.Handle).Methods("GET")
	app.Router.HandleFunc("/", getRequest).Methods("GET")
	app.Router.HandleFunc("/", postRequest).Methods("POST")
	app.Router.HandleFunc("/", deleteRequest).Methods("DELETE")

	app.Router.HandleFunc("/products", app.AllProducts).Methods("GET")
	app.Router.HandleFunc("/products/{id}", app.FetchProduct).Methods("GET")
	app.Router.HandleFunc("/products", app.NewProducts).Methods("POST")
}

func (app *App) Run() {
	http.Handle("/", app.Router)
	log.Printf("Server started and listening on localhost%s\n", app.Port)

	err := http.ListenAndServe(app.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
