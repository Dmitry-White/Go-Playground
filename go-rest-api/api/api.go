package api

import (
	"database/sql"
	"go-rest-api/api/data"
	"go-rest-api/api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	*data.App
}

func (app *App) Init() {
	db, err := sql.Open("sqlite3", app.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()

	app.DB = db
	app.Router = router
	app.initRoutes()
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/handle", handlers.Handle(app.App)).Methods("GET")
	app.Router.HandleFunc("/", handlers.GetRequest).Methods("GET")
	app.Router.HandleFunc("/", handlers.PostRequest).Methods("POST")
	app.Router.HandleFunc("/", handlers.DeleteRequest).Methods("DELETE")

	app.Router.HandleFunc("/products", handlers.AllProducts(app.App)).Methods("GET")
	app.Router.HandleFunc("/products", handlers.NewProduct(app.App)).Methods("POST")
	app.Router.HandleFunc("/products/{id}", handlers.FetchProduct(app.App)).Methods("GET")

	app.Router.HandleFunc("/orders", handlers.AllOrders(app.App)).Methods("GET")
	app.Router.HandleFunc("/orders", handlers.NewOrder(app.App)).Methods("POST")
	app.Router.HandleFunc("/orders/{id}", handlers.FetchOrder(app.App)).Methods("GET")
	app.Router.HandleFunc("/orders/{id}/items", handlers.FetchOrderItems(app.App)).Methods("GET")

	app.Router.HandleFunc("/order-items", handlers.AllOrderItems(app.App)).Methods("GET")
	app.Router.HandleFunc("/order-items", handlers.NewOrderItem(app.App)).Methods("POST")
}

func (app *App) Run() {
	http.Handle("/", app.Router)
	log.Printf("Server started and listening on localhost%s\n", app.Port)

	err := http.ListenAndServe(app.Port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
