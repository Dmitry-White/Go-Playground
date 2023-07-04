package main

import (
	"fmt"
	"text/template"
)

func render(t string) (*template.Template, error) {
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
		return nil, err
	}

	return tmpl, nil
}
