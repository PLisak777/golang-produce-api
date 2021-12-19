package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/PLisak777/golang-produce-api/pkg/mocks"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock items
	for _, produce := range mocks.Groceries {
		if produce.ID == id {
			// If IDs are equal send item s a response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			r.Header.Add("Access-Control-Allow-Origin", "*")

			json.NewEncoder(w).Encode(produce)
			break
		}
	}
}