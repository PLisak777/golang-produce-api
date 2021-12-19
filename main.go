package main

import (
	"log"
	"net/http"

	"github.com/PLisak777/golang-produce-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/produce", handlers.GetAllItems).Methods(http.MethodGet)
	router.HandleFunc("/produce", handlers.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/produce/{id}", handlers.GetById).Methods(http.MethodGet)
	router.HandleFunc("/produce/{id}", handlers.DeleteItem).Methods(http.MethodDelete)

	log.Println("API Running")
	http.ListenAndServe(":4000", router)
}