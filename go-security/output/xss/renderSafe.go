package main

import (
	"go-security/output/xss/dal"
	"go-security/output/xss/pages"
	"html/template"
	"net/http"
)

func renderSafe(resw http.ResponseWriter, ms []dal.Message, user string) error {
	params := dal.Params{
		Name:     user,
		Messages: ms,
	}
	htmlTemplate := template.Must(template.New("Safe").Parse(pages.TemplateSafe))

	err := htmlTemplate.Execute(resw, params)
	if err != nil {
		return err
	}

	return nil
}
