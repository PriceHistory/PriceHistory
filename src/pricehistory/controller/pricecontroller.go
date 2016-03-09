package controller

import (
	"github.com/ungerik/go-dry"
	"encoding/json"
	"net/http"
	"pricehistory/database"
)

func GetProductWithPrices(w http.ResponseWriter, r *http.Request) {
	productOuterID := r.URL.Query().Get("id")
	product := database.GetProductWithPrices(productOuterID)
	jsonResponse, err := json.Marshal(product)
	dry.PanicIfErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func Run() {
	http.HandleFunc("/prices", GetProductWithPrices)
	http.ListenAndServe(":8080", nil)
}
