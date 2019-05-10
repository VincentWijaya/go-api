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

	log.Fatal(http.ListenAndServe(":3000", r))
}
