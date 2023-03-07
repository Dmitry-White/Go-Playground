package tests

import (
	"encoding/json"
	"go-rest-api/api"
	"net/http"
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

func TestGetNonExistentProduct(t *testing.T) {
	clearProductTable(&app)

	request, _ := http.NewRequest("GET", "/product/13", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	m := map[string]string{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf(
			"Expected the 'error' key of the response to be set to 'sql: no rows in result set'. Got '%s'",
			m["error"],
		)
	}
}
