package handlers

import (
	"encoding/json"
	"golang-produce-api/pkg/mocks"
	"net/http"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r.Header.Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(mocks.Groceries)
}
