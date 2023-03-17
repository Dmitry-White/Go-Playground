package handlers

import (
	"context"
	"net/http"
	"sync"
	"time"

	"go-applied-concurrency/api/db"
	"go-applied-concurrency/api/repo"
	"go-applied-concurrency/api/services"

	"github.com/gorilla/mux"
)

type Handler struct {
	repo  repo.IRepo
	stats services.IStatistics
	once  sync.Once
}

type IHandler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Close(w http.ResponseWriter, r *http.Request)
	Stats(w http.ResponseWriter, r *http.Request)
	ProductIndex(w http.ResponseWriter, r *http.Request)
	OrderShow(w http.ResponseWriter, r *http.Request)
	OrderInsert(w http.ResponseWriter, r *http.Request)
	OrderReverse(w http.ResponseWriter, r *http.Request)
}

func New() (IHandler, error) {
	r, err := repo.New(db.ProductInputPath)
	if err != nil {
		return nil, err
	}

	_, processed, done := r.Index()
	s := services.New(processed, done)

	h := Handler{
		repo:  r,
		stats: s,
	}
	return &h, nil
}

// Index returns a simple hello response for the homepage
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	// Send an HTTP status & a hardcoded message
	h.repo.Index()
	writeResponse(w, http.StatusOK, "Welcome to the Orders App!", nil)
}

func (h *Handler) Close(w http.ResponseWriter, r *http.Request) {
	h.once.Do(func() {
		h.repo.Close()
	})
	writeResponse(w, http.StatusOK, "The Orders App is now closed!", nil)
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	ctx, cancel := context.WithTimeout(reqCtx, 500*time.Millisecond)
	defer cancel()

	stats, err := h.stats.GetStats(ctx)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, nil, err)
		return
	}

	writeResponse(w, http.StatusOK, stats, nil)
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
	router.Methods("DELETE").Path("/orders/{orderId}").
		Handler(http.HandlerFunc(handler.OrderReverse))
	router.Methods("POST").Path("/orders").
		Handler(http.HandlerFunc(handler.OrderInsert))
	router.Methods("POST").Path("/close").
		Handler(http.HandlerFunc(handler.Close))
	router.Methods("GET").Path("/stats").
		Handler(http.HandlerFunc(handler.Stats))

	return router
}
