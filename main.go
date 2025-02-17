package main

import (
	"log"
	"net/http"

	"github.com/cloudarmin/review/handlers"
)

func main() {
	bookHandler := handlers.NewBookHandler()
	http.HandleFunc("/books", bookHandler.GetBooks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
