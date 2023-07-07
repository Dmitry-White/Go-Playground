package main

import (
	"fmt"
	"text/template"
)

func render(t string) (*template.Template, error) {
	templateSlice := []string{}
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", PATHS.PAGES, t))

	partials := []string{
		PATHS.TEMPLATES + "/base.layout.gohtml",
		PATHS.TEMPLATES + "/buttons.partial.gohtml",
		PATHS.TEMPLATES + "/header.partial.gohtml",
		PATHS.TEMPLATES + "/footer.partial.gohtml",
	}
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
