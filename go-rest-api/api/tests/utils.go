package tests

import (
	"go-rest-api/api"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ensureTableExists(app *api.App) {
	if _, err := app.DB.Exec(tableProductCreationQuery); err != nil {
		log.Fatalln(err.Error())
	}
}

func clearProductTable(app *api.App) {
	if _, err := app.DB.Exec(tableProductClearingQuery); err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := app.DB.Exec(tableProductDeletionQuery); err != nil {
		log.Fatalln(err.Error())
	}
}

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
