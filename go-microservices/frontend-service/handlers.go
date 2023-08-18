package main

import (
	"net/http"
	"os"
)

func handleIndex(resw http.ResponseWriter, r *http.Request) {
	tmpl, err := render("index.gohtml")
	if err != nil {
		http.Error(resw, err.Error(), http.StatusInternalServerError)
	}

	data := Data{
		Mode: os.Getenv("MODE"),
	}

	err = tmpl.Execute(resw, data)
	if err != nil {
		http.Error(resw, err.Error(), http.StatusInternalServerError)
	}
}

var handleJS = http.StripPrefix(ROUTES.JS, http.FileServer(http.Dir(PATHS.JS)))
