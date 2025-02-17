package handlers

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookHandler struct {
	books []Book
}

func NewBookHandler() *BookHandler {
	return &BookHandler{
		books: []Book{
			{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
			{ID: "2", Title: "Clean Code", Author: "Robert C. Martin"},
		},
	}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.books)
}

func (h *BookHandler) GetBookByID(id string) (Book, bool) {
	for _, book := range h.books {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}
