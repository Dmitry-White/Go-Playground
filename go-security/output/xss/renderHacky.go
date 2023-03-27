package main

import (
	"fmt"
	"go-security/output/xss/dal"
	"go-security/output/xss/pages"
	"net/http"
)

func renderHacky(resw http.ResponseWriter, ms []dal.Message, user string) error {
	body := dal.FormateMessages(ms)

	_, err := fmt.Fprintf(resw, pages.TemplateHacky, user, user, len(ms), body)
	if err != nil {
		return err
	}

	return nil
}
