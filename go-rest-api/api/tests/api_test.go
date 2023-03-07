package tests

import (
	"go-rest-api/api"
	"os"
	"testing"
)

var app api.App

func TestMain(m *testing.M) {
	app = api.App{}
	app.Init()
	ensureTableExists(&app)
	code := m.Run()

	clearProductTable(&app)
	os.Exit(code)
}
