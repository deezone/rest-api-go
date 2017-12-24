package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/quote", GetQuotes).Methods("GET")
	router.HandleFunc("/quote/{id}", GetQuote).Methods("GET")
	router.HandleFunc("/quote/{id}", CreateQuote).Methods("POST")
	router.HandleFunc("/quote/{id}", DeleteQuote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", router))
}

// GetQuotes looks up all of the quotes.
func GetQuotes(w http.ResponseWriter, r *http.Request) {}

// GetQuote looks up a specific quote by ID.
func GetQuote(w http.ResponseWriter, r *http.Request) {}

// CreateQuote creates a new quote.
func CreateQuote(w http.ResponseWriter, r *http.Request) {}

// DeleteQuote deletes a quote by quote ID.
func DeleteQuote(w http.ResponseWriter, r *http.Request) {}
