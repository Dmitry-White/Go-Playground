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

func AllProducts(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("AllProducts Request: %+v\n", r)

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
		log.Printf("FetchProduct Request: %+v\n", r)
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		p := data.Product{}
		p.ID = id
		err := services.GetProduct(app.DB, &p)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, p)
	}
}

func NewProducts(app *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("NewProducts Request: %+v\n", r)

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		var p data.Product
		err = json.Unmarshal(reqBody, &p)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		err = services.CreateProduct(app.DB, &p)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, p)
	}
}
