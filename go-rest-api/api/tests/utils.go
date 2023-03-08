package tests

import (
	"go-rest-api/api"
	"go-rest-api/api/services"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ensureTableExists(tableCreationQuery string) func(*api.App) {
	return func(app *api.App) {
		if _, err := app.DB.Exec(tableCreationQuery); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

var ensureProductTableExists = ensureTableExists(productTableCreationQuery)
var ensureOrderTableExists = ensureTableExists(orderTableCreationQuery)
var ensureOrderItemTableExists = ensureTableExists(orderItemTableCreationQuery)

func ensureItemExists(creationQuery string) func(*api.App) {
	return func(app *api.App) {
		if _, err := app.DB.Exec(creationQuery); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

var ensureProductExists = ensureItemExists(productCreationQuery)
var ensureOrderExists = ensureItemExists(orderCreationQuery)
var ensureOrderItemExists = ensureItemExists(orderItemCreationQuery)

func clearTable(tableClearingQuery, tableDeletionQuery string) func(*api.App) {
	return func(app *api.App) {
		if _, err := app.DB.Exec(tableClearingQuery); err != nil {
			log.Fatalln(err.Error())
		}

		if _, err := app.DB.Exec(tableDeletionQuery); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

var clearProductTable = clearTable(productTableClearingQuery, productTableDeletionQuery)
var clearOrderTable = clearTable(orderTableClearingQuery, orderTableDeletionQuery)
var clearOrderItemTable = clearTable(orderItemTableClearingQuery, orderItemTableDeletionQuery)

func executeRequest(app *api.App, req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, req)

	return recorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseField(t *testing.T, field, expected, actual interface{}) {
	if actual != expected {
		t.Errorf(
			"Expected the '%s' to be set to '%s'. Got '%s'",
			field,
			expected,
			actual,
		)
	}
}

func checkTables(app *api.App) {
	services.GetProducts(app.DB)
	services.GetOrders(app.DB)
	services.GetItems(app.DB)
}
