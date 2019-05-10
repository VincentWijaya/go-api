package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vincentwijaya/go-api/controllers"
)

func main() {
	// Declare Router
	r := mux.NewRouter()

	// Declare controllers
	book := controllers.NewBookController()

	// Endpoints
	r.HandleFunc("/api/books", book.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", book.GetBook).Methods("GET")

	log.Printf("Running on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
