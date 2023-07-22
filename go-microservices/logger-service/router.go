package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *AppConfig) router() *chi.Mux {
	mux := chi.NewRouter()
	return mux
}
