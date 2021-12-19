package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang-produce-api/pkg/mocks"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, produce := range mocks.Groceries {
		if produce.ID == id {
			// Delete item and send resp if ID matches dynamic ID
			mocks.Groceries = append(mocks.Groceries[:index], mocks.Groceries[index+1:]...)

			w.Header().Add("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			r.Header.Add("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}