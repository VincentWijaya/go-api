package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vincentwijaya/go-api/models"
)

// BookController struct
type BookController struct{}

var books []models.Book

func init() {
	// Mock Data
	books = append(books, models.Book{ID: "1", Isbn: "453636", Title: "Book One", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "2", Isbn: "3434343", Title: "Book Two", Author: &models.Author{Firstname: "Steve", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "3", Isbn: "123123", Title: "Book Three", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
}

// NewBookController init
func NewBookController() *BookController {
	return &BookController{}
}

// GetBooks func
func (c BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func (c BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Loop through books and find the id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Book{})
}
