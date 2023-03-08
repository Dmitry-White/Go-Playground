package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app.Init()
	ensureTableExists(&app)
	code := m.Run()

	clearProductTable(&app)
	os.Exit(code)
}
