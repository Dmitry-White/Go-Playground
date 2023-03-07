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
		log.Fatalln(err)
	}
}

func clearProductTable(app *api.App) {
	if _, err := app.DB.Exec(tableProductClearingQuery); err != nil {
		log.Fatalln(err)
	}

	if _, err := app.DB.Exec(tableProductDeletionQuery); err != nil {
		log.Fatalln(err)
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
