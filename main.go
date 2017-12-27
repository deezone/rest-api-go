// A learning excercise to understand the use of the Go programming language when building an API
// Started Christmas 2017
// Governed by the license that can be found in the LICENSE file
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"time"
	"strconv"
	"github.com/tkanos/gonfig"
	"strings"
	"os"
)

// Configuration type settings common to application
type Configuration struct {
	Port        int
	Environment string
}

// The Quote type (more like an object) that manages the details of a quote.
type Quote struct {
	ID      int      `json:"id"`
	Quote   string   `json:"quote"`
	Authour *Authour `json:"authour,omitempty"`
}

// Authour type, referenced by core items: quotes, publications, etc.
type Authour struct {
	ID          int       `json:"id"`
	First       string    `json:"first,omitempty"`
	Last        string    `json:"last,omitempty"`
	Born        time.Time `json:"born,omitempty"`
	Died        time.Time `json:"died,omitempty"`
	Description string    `json:"description,omitempty"`
	BioLink     string    `json:"biolink,omitempty"`
}

var quotes []Quote

// GetQuotes looks up all of the quotes.
// GET /quotes
// Returns all of the quotes in the JSON format.
func GetQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

// GetQuote looks up a specific quote by ID.
// GET /quote/{id}
// Returns a quote in the JSON format provided the target ID is valid.
func GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		io.WriteString(w, `{"status": "error, invalid ID"}`)
		return
	}

	for _, item := range quotes {
		if item.ID == targetID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Quote{})
}

// CreateQuote creates a new quote.
// POST /quote/{id}
// Returns all quotes including the newly added quote made by the POST request.
func CreateQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		io.WriteString(w, `{"status": "error, invalid ID"}`)
		return
	}

	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)
	quote.ID = targetID
	quotes = append(quotes, quote)
	json.NewEncoder(w).Encode(quotes)
}

// DeleteQuote deletes a quote by quote ID.
// DELETE /quote/{id}
// Returns all quotes which will exclude the deleted quote made by the DELETE request.
func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	targetID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		io.WriteString(w, `{"status": "error, invalid ID"}`)
		return
	}

	for index, item := range quotes {
		if item.ID == targetID {
			quotes = append(quotes[:index], quotes[:index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(quotes)
}

// main function
func main() {
	configuration := Configuration{}
	configuration.Environment = os.Getenv("REST_API_ENV")
	env := []string{}
	env = append(env, "config/config.", configuration.Environment, ".json")
	err := gonfig.GetConf(strings.Join(env, ""), &configuration)
	if (err != nil) { return }

	quotes = LoadData()

	router := mux.NewRouter()

	router.HandleFunc("/quotes", GetQuotes).Methods("GET")
	router.HandleFunc("/quote/{id}", GetQuote).Methods("GET")
	router.HandleFunc("/quote/{id}", CreateQuote).Methods("POST")
	router.HandleFunc("/quote/{id}", DeleteQuote).Methods("DELETE")

	port := []string{}
	port = append(port, ":", strconv.Itoa(configuration.Port))
	log.Fatal(http.ListenAndServe(strings.Join(port, ""), router))
}

// LoadData manages the gathering of quote data.
// Used as data store to mimic a database
// Returns an array of quotes
func LoadData() []Quote {
	var data []Quote

	data = append(data, Quote{
		ID: 1,
		Quote: "When it's pouring rain and you're bowling along through the wet, there's satisfaction in knowing you're out there and the others aren't.",
		Authour: &Authour{
			ID: 1,
			First: "Peter",
			Last: "Snell",
			Born: time.Date(1938, time.December, 17, 0, 0, 0, 0, time.UTC),
			Description: "A New Zealand former middle-distance runner. He won three Olympic gold medals, and is the only male since 1920 to win the 800 and 1500 metres at the same Olympics, in 1964.",
			BioLink:"https://en.wikipedia.org/wiki/Peter_Snell"}})
	data = append(data, Quote{
		ID: 2,
		Quote: "I run because it's my passion, and not just a sport. Every time I walk out the door, I know why I'm going where I'm going and I'm already focused on that special place where I find my peace and solitude. Running, to me, is more than just a physical exercise... it's a consistent reward for victory!",
		Authour: &Authour{
			ID: 2,
			First: "Sasha",
			Last: "Azevedo",
			Born: time.Date(1978, time.May, 20, 0, 0, 0, 0, time.UTC),
			Description: "Modeling, acting and photography",
			BioLink:"http://www.imdb.com/name/nm1659315/bio"}})
	data = append(data, Quote{
		ID: 3,
		Quote: "If you always put limits on what you can do, physical or anything else, it'll spread over into the rest of your life. It'll spread into your work, into your morality, into your entire being. There are no limits. There are plateaus, but you must not stay there, you must go beyond them.",
		Authour: &Authour{
			ID: 3,
			First: "Bruce",
			Last: "Lee",
			Born: time.Date(1940, time.November, 27, 0, 0, 0, 0, time.UTC),
			Died: time.Date(1973, time.July, 20, 0, 0, 0, 0, time.UTC),
			Description: "A Hong Kong and American actor, film director, martial artist, martial arts instructor, philosopher[5] and founder of the martial art Jeet Kune Do.",
			BioLink:"https://en.wikipedia.org/wiki/Bruce_Lee"}})
	data = append(data, Quote{
		ID: 4,
		Quote: "Motivation is what gets you started. Habit is what keeps you going.",
		Authour: &Authour{
			ID: 4,
			First: "Jim",
			Last: "Ruyn",
			Born: time.Date(1947, time.April, 29, 0, 0, 0, 0, time.UTC),
			Description: "A former American politician and track and field athlete. He won a silver medal in the men's 1500 metres at the 1968 Summer Olympics, and was the first high school athlete to run a mile in under four minutes.",
			BioLink:"https://en.wikipedia.org/wiki/Jim_Ryun"}})
	data = append(data, Quote{
		ID: 5,
		Quote: "I don't think about the miles that are coming down the road, I don't think about the mile I'm on right now, I don't think about the miles I've already covered. I think about what I'm doing right now, just being lost in the moment.",
		Authour: &Authour{
			ID: 5,
			First: "Ryan",
			Last: "Hall",
			Born: time.Date(1982, time.October, 14, 0, 0, 0, 0, time.UTC),
			Description: "U.S. Olympic marathoner",
			BioLink:"https://en.wikipedia.org/wiki/Ryan_Hall_(runner)"}})

	return data
}
