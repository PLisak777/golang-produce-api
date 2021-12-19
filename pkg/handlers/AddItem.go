package handlers

import (
	"encoding/json"
	"golang-produce-api/pkg/mocks"
	"golang-produce-api/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var produce models.Produce
	json.Unmarshal(body, &produce)

	// Append to the Item mocks
	produce.ID = rand.Intn(100)
	mocks.Groceries = append(mocks.Groceries, produce)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	r.Header.Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode("Created")
}
