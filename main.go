package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"time"
)

// The quote Type (more like an object)
type Quote struct {
	ID      int      `json:"id"`
	Quote   string   `json:"quote"`
	Authour *Authour `json:"authour,omitempty"`
}
type Authour struct {
	First     string    `json:"first,omitempty"`
	Last      string    `json:"last,omitempty"`
	Born      time.Time `json:"born,omitempty"`
	Died      time.Time `json:"died,omitempty"`
	Wikipedia string    `json:"wikipedia,omitempty"`
}

var quotes []Quote

// GetQuotes looks up all of the quotes.
func GetQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

// GetQuote looks up a specific quote by ID.
func GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

// CreateQuote creates a new quote.
func CreateQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

// DeleteQuote deletes a quote by quote ID.
func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

// main function
func main() {
	router := mux.NewRouter()

	quotes = append(quotes, Quote{
		ID: "1",
		Quote: "When it's pouring rain and you're bowling along through the wet, there's satisfaction in knowing you're out there and the others aren't.",
		Authour: &Authour{
			First: "Peter",
			Last: "Snell",
			Born: "17 December 1938",
			Died: "",
			Wikipedia:"https://en.wikipedia.org/wiki/Peter_Snell"}})


	router.HandleFunc("/quotes", GetQuotes).Methods("GET")
	router.HandleFunc("/quote/{id}", GetQuote).Methods("GET")
	router.HandleFunc("/quote/{id}", CreateQuote).Methods("POST")
	router.HandleFunc("/quote/{id}", DeleteQuote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", router))
}
