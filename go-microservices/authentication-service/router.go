package main

import "github.com/go-chi/chi/v5"

func (a *AppConfig) router() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(loggerMiddleware)
	mux.Use(healthMiddleware)
	mux.Use(corsMiddleware)

	mux.Post(ROUTES.INDEX, a.handleIndex)

	return mux
}
