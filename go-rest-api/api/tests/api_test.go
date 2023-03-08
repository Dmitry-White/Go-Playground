package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app.Init()
	ensureProductTableExists(&app)
	ensureOrderTableExists(&app)
	code := m.Run()

	clearProductTable(&app)
	clearOrderTable(&app)
	os.Exit(code)
}
