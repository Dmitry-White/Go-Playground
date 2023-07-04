package main

import "net/http"

func handleIndex(resw http.ResponseWriter, r *http.Request) {
	tmpl, err := render("index.gohtml")

	if err != nil {
		http.Error(resw, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(resw, nil)
	if err != nil {
		http.Error(resw, err.Error(), http.StatusInternalServerError)
	}
}
