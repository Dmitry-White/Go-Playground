package handlers

import (
	"encoding/json"
	"go-rest-api/api/data"
	"go-rest-api/api/services"
	"go-rest-api/api/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllProducts(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := services.GetProducts(app.DB)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, products)
	}
}

func FetchProduct(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		product := data.Product{
			ID: id,
		}
		err := services.GetProduct(app.DB, &product)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, product)
	}
}

func NewProduct(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		product := data.Product{}
		err = json.Unmarshal(reqBody, &product)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		err = services.CreateProduct(app.DB, &product)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, product)
	}
}
