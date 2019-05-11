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
	r.HandleFunc("/api/books", book.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", book.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", book.DeleteBook).Methods("DELETE")

	log.Printf("Running on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
