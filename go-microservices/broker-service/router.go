package main

import "github.com/go-chi/chi/v5"

func (app *AppConfig) router() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(loggerMiddleware)
	mux.Use(healthMiddleware)
	mux.Use(corsMiddleware)

	mux.Post(ROUTES.INDEX, app.handleIndex)
	mux.Post(ROUTES.PROCESS, app.handleProcess)

	return mux
}
