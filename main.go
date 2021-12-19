package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang-produce-api/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/produce", handlers.GetAllItems).Methods(http.MethodGet)
	router.HandleFunc("/items", handlers.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", handlers.GetById).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods(http.MethodDelete)

	log.Println("API Running")
	http.ListenAndServe(":4000", router)
}