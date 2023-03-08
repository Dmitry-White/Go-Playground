package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestGetNonExistentOrderItems(t *testing.T) {
	clearOrderTable(&app)
	clearOrderItemTable(&app)

	request, _ := http.NewRequest("GET", "/orders/13/items", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	m := []map[string]int{}
	json.Unmarshal(response.Body.Bytes(), &m)

	expected := 0
	actual := len(m)
	checkResponseField(t, "length", expected, actual)
}

func TestCreateOrderItem(t *testing.T) {
	clearOrderTable(&app)
	clearOrderItemTable(&app)

	checkTables(&app)

	payload := []byte(`{
		"orderId": 1,
		"productId": 1,
		"quantity": 1
	}`)
	request, _ := http.NewRequest("POST", "/order-items", bytes.NewBuffer(payload))
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]int
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		log.Fatal(err.Error())
	}

	checkResponseField(t, "orderId", 1, m["orderId"])
	checkResponseField(t, "productId", 1, m["productId"])
	checkResponseField(t, "quantity", 1, m["quantity"])
}

func TestGetOrderItems(t *testing.T) {
	clearOrderTable(&app)
	clearOrderItemTable(&app)

	ensureOrderExists(&app)
	ensureOrderItemExists(&app)

	request, _ := http.NewRequest("GET", "/orders/1/items", nil)
	response := executeRequest(&app, request)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m [](map[string]int)
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseField(t, "orderId", 1, m[0]["orderId"])
	checkResponseField(t, "productId", 1, m[0]["productId"])
	checkResponseField(t, "quantity", 1, m[0]["quantity"])
}
