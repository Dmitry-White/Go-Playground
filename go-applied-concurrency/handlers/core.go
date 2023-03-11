package handlers

import (
	"net/http"

	"go-applied-concurrency/db"
	"go-applied-concurrency/repo"

	"github.com/gorilla/mux"
)

type Handler struct {
	repo repo.IRepo
}

type IHandler interface {
	Index(w http.ResponseWriter, r *http.Request)
	ProductIndex(w http.ResponseWriter, r *http.Request)
	OrderShow(w http.ResponseWriter, r *http.Request)
	OrderInsert(w http.ResponseWriter, r *http.Request)
}

func New() (IHandler, error) {
	r, err := repo.New(db.ProductInputPath)
	if err != nil {
		return nil, err
	}
	h := Handler{repo: r}
	return &h, nil
}

// Index returns a simple hello response for the homepage
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP status & a hardcoded message
	writeResponse(w, http.StatusOK, "Welcome to the Orders App!", nil)
}

// InitRoutes configures the routes of this handler and binds handler functions to them
func InitRoutes(handler IHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").
		Handler(http.HandlerFunc(handler.Index))
	router.Methods("GET").Path("/products").
		Handler(http.HandlerFunc(handler.ProductIndex))
	router.Methods("GET").Path("/orders/{orderId}").
		Handler(http.HandlerFunc(handler.OrderShow))
	router.Methods("POST").Path("/orders").
		Handler(http.HandlerFunc(handler.OrderInsert))

	return router
}
