package handlers

import (
	"encoding/json"
	"go-rest-api/api/data"
	"go-rest-api/api/services"
	"go-rest-api/api/utils"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllOrderItems(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("AllOrderItems Request: %+v\n", r)

		orderItems, err := services.GetOrderItems(app.DB)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, orderItems)
	}
}

func FetchOrderItem(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("FetchOrderItem Request: %+v\n", r)
		params := mux.Vars(r)
		query := r.URL.Query()

		orderId, _ := strconv.Atoi(params["orderId"])
		productId, _ := strconv.Atoi(query.Get("productId"))

		orderItems := data.OrderItem{}
		orderItems.OrderId = orderId
		orderItems.ProductId = productId
		err := services.GetOrderItem(app.DB, &orderItems)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, orderItems)
	}
}

func NewOrderItem(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("NewOrderItem Request: %+v\n", r)

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		orderItem := data.OrderItem{}
		err = json.Unmarshal(reqBody, &orderItem)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		err = services.CreateOrderItem(app.DB, &orderItem)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, orderItem)
	}
}
