package main

import (
	"net/http"
)

func (a *AppConfig) router() *http.ServeMux {
	router := http.ServeMux{}

	router.HandleFunc(ROUTES.INDEX, handleIndex)
	router.Handle(ROUTES.JS, handleJS)

	return &router
}
