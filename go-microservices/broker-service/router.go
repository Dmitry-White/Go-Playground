package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *Config) router() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Get(ROUTES.INDEX, handleIndex)

	return mux
}
