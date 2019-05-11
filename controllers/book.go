package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

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

// GetBook func
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

// CreateBook func
func (c BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock Id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook func
func (c BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Loop through books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(books)
}

// DeleteBook func
func (c BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Loop through books and find the id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}
