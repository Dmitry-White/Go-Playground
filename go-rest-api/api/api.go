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
	app.Router.HandleFunc("/handle", handlers.Handle(app.App)).Methods("GET")
	app.Router.HandleFunc("/", handlers.GetRequest).Methods("GET")
	app.Router.HandleFunc("/", handlers.PostRequest).Methods("POST")
	app.Router.HandleFunc("/", handlers.DeleteRequest).Methods("DELETE")

	app.Router.HandleFunc("/products", handlers.AllProducts(app.App)).Methods("GET")
	app.Router.HandleFunc("/products/{id}", handlers.FetchProduct(app.App)).Methods("GET")
	app.Router.HandleFunc("/products", handlers.NewProduct(app.App)).Methods("POST")

	app.Router.HandleFunc("/orders", handlers.AllOrders(app.App)).Methods("GET")
	app.Router.HandleFunc("/orders/{id}", handlers.FetchOrder(app.App)).Methods("GET")
	app.Router.HandleFunc("/orders", handlers.NewOrder(app.App)).Methods("POST")

	app.Router.HandleFunc("/order-items", handlers.AllOrderItems(app.App)).Methods("GET")
	app.Router.HandleFunc("/order-items/{orderId}", handlers.FetchOrderItem(app.App)).Methods("GET")
	app.Router.HandleFunc("/order-items", handlers.NewOrderItem(app.App)).Methods("POST")
}

func (app *App) Run() {
	http.Handle("/", app.Router)
	log.Printf("Server started and listening on localhost%s\n", app.Port)

	err := http.ListenAndServe(app.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
