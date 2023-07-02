package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	render(w, "index.gohtml")
}

func main() {
	http.HandleFunc(ROUTES.INDEX, handleIndex)

	log.Printf("Server listening on %s", ADDR)

	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	templateSlice := []string{}
	templateSlice = append(templateSlice, fmt.Sprintf("./web/pages/%s", t))

	partials := []string{
		"./web/templates/base.layout.gohtml",
		"./web/templates/header.partial.gohtml",
		"./web/templates/footer.partial.gohtml",
	}
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
