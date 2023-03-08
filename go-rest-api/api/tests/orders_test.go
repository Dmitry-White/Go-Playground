package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetNonExistentOrder(t *testing.T) {
	clearOrderTable(&app)

	request, _ := http.NewRequest("GET", "/orders/13", nil)
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

func TestCreateOrder(t *testing.T) {
	clearOrderTable(&app)
	clearOrderItemTable(&app)

	checkTables(&app)

	payload := []byte(`{
		"customerName": "Mikki Mouse",
		"total": 200,
		"status": "testing",
		"items": [
			{
				"productId": 1,
				"quantity": 1
			}
		]
	}`)
	request, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(payload))
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseField(t, "customerName", "Mikki Mouse", m["customerName"])
	checkResponseField(t, "total", 200.0, m["total"])
	checkResponseField(t, "status", "testing", m["status"])
	checkResponseField(t, "id", 1.0, m["id"])
}

func TestGetOrder(t *testing.T) {
	clearOrderTable(&app)

	ensureOrderExists(&app)

	request, _ := http.NewRequest("GET", "/orders/1", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseField(t, "customerName", "Mikki Mouse", m["customerName"])
	checkResponseField(t, "total", 200.0, m["total"])
	checkResponseField(t, "status", "testing", m["status"])
	checkResponseField(t, "id", 1.0, m["id"])
}
