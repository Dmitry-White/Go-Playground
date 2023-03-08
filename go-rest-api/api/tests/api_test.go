package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app.Init()
	ensureProductTableExists(&app)
	ensureOrderTableExists(&app)
	ensureOrderItemTableExists(&app)
	code := m.Run()

	clearProductTable(&app)
	clearOrderTable(&app)
	clearOrderItemTable(&app)
	os.Exit(code)
}
