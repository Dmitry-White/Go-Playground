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

func AllOrders(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("AllOrders Request: %+v\n", r)

		orders, err := services.GetOrders(app.DB)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, orders)
	}
}

func FetchOrder(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("FetchOrder Request: %+v\n", r)
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		order := data.Order{
			Id: id,
		}
		err := services.GetOrder(app.DB, &order)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, order)
	}
}

func NewOrder(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("NewOrder Request: %+v\n", r)

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		order := data.Order{}
		err = json.Unmarshal(reqBody, &order)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		err = services.CreateOrder(app.DB, &order)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		for _, item := range order.Items {
			orderItem := data.OrderItem{}
			orderItem = item
			orderItem.OrderId = order.Id
			err = services.CreateOrderItem(app.DB, &orderItem)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, err)
				return
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, order)
	}
}
