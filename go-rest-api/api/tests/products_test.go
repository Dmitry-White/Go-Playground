package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetNonExistentProduct(t *testing.T) {
	clearProductTable(&app)

	request, _ := http.NewRequest("GET", "/products/13", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	m := map[string]string{}
	json.Unmarshal(response.Body.Bytes(), &m)

	expected := "sql: no rows in result set"
	actual := m["error"]
	if actual != expected {
		t.Errorf(
			"Expected the 'error' to be set to '%s'. Got '%s'",
			expected,
			actual,
		)
	}
}

func TestCreateProduct(t *testing.T) {
	clearProductTable(&app)

	checkTables(&app)

	payload := []byte(`{
		"productCode": "TEST12345",
		"name": "ProductTest",
		"inventory": 1,
		"price": 1,
		"status": "testing"
	}`)
	request, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseField(t, "productCode", "TEST12345", m["productCode"])
	checkResponseField(t, "name", "ProductTest", m["name"])
	checkResponseField(t, "inventory", 1.0, m["inventory"])
	checkResponseField(t, "price", 1.0, m["price"])
	checkResponseField(t, "status", "testing", m["status"])
	checkResponseField(t, "id", 1.0, m["id"])
}

func TestGetProduct(t *testing.T) {
	clearProductTable(&app)

	ensureProductExists(&app)

	request, _ := http.NewRequest("GET", "/products/1", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseField(t, "productCode", "TEST12345", m["productCode"])
	checkResponseField(t, "name", "ProductTest", m["name"])
	checkResponseField(t, "inventory", 1.0, m["inventory"])
	checkResponseField(t, "price", 1.0, m["price"])
	checkResponseField(t, "status", "testing", m["status"])
	checkResponseField(t, "id", 1.0, m["id"])
}
