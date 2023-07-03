package main

import (
	"github.com/go-chi/chi/v5"
)

func (a *Config) router() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(loggerMiddleware)
	mux.Use(healthMiddleware)
	mux.Use(corsMiddleware)

	mux.Get(ROUTES.INDEX, handleIndex)

	return mux
}
