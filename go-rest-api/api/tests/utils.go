package tests

import (
	"go-rest-api/api"
	"go-rest-api/api/data"
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

func ensureItemExists(creationQuery string) func(*api.App) {
	return func(app *api.App) {
		if _, err := app.DB.Exec(creationQuery); err != nil {
			log.Fatalln(err.Error())
		}

		rows, err := app.DB.Query("SELECT * FROM orders")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer rows.Close()

		orders := []data.Order{}

		for rows.Next() {
			order := data.Order{}
			err := rows.Scan(&order.Id, &order.CustomerName, &order.Total, &order.Status)
			if err != nil {
				log.Fatalln(err.Error())
			}

			orders = append(orders, order)
		}
		log.Printf("Orders: %+v\n", orders)
	}
}

var ensureProductExists = ensureItemExists(productCreationQuery)
var ensureOrderExists = ensureItemExists(orderCreationQuery)

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
