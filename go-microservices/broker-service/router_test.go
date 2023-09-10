package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestRouter(t *testing.T) {
	app := &AppConfig{}

	actual := app.router().Routes()
	expected := []string{ROUTES.INDEX, ROUTES.PROCESS}
	errors := []string{}

	for _, actualRoute := range actual {
		if !slices.Contains(expected, actualRoute.Pattern) {
			errors = append(errors, actualRoute.Pattern)
		}
	}

	if len(actual) == 0 {
		t.Error("No routes registered!")
	}

	if len(errors) != 0 {
		t.Error("Following routes are not expected!", errors)
	}
}
