package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	handler := NewBookHandler()

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetBooks(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var books []Book
	err = json.NewDecoder(rr.Body).Decode(&books)
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}
}

func TestGetBookByID(t *testing.T) {
	handler := NewBookHandler()

	book, found := handler.GetBookByID("1")
	if !found {
		t.Error("expected to find book with ID 1")
	}

	if book.Title != "The Go Programming Language" {
		t.Errorf("expected book title 'The Go Programming Language', got '%s'", book.Title)
	}

	_, found = handler.GetBookByID("999")
	if found {
		t.Error("expected not to find book with ID 999")
	}
}
