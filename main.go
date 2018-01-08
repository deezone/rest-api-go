// A learning excercise to understand the use of the Go programming language when building an API
// Started Christmas 2017
// Governed by the license that can be found in the LICENSE file
package main

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"strconv"
	"strings"
	"os"
	"fmt"
	"runtime"

	"github.com/tkanos/gonfig"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Configuration type, settings common to application
type Configuration struct {
	Port        int
	Version     string
	ReleaseDate string
	Environment string
	DbHost      string
	DbUser      string
	DbName      string
	DbPassword  string
}

// Quote type (more like an object), manages the details of a quote.
type Quote struct {
	gorm.Model

	Quote    string  `json:"quote"`
	AuthorID uint    `json:"author-id"`
}

// Author type, referenced by core items: quotes, publications, etc.
type Author struct {
	gorm.Model

	First       string    `json:"first,omitempty"`
	Last        string    `json:"last,omitempty"`
	Born        time.Time `json:"born,omitempty"`
	Died        time.Time `json:"died,omitempty"`
	Description string    `json:"description,omitempty"`
	BioLink     string    `json:"biolink,omitempty"`
	Quotes      []Quote   `gorm:"ForeignKey:AuthorID";json:"quotes,omitempty"`
}

type Health struct {
	Refer      string `json:"reference,omitempty"`
	Alloc      uint64 `json:"alloc,omitempty"`
	TotalAlloc uint64 `json:"total-alloc,omitempty"`
	Sys        uint64 `json:"sys,omitempty"`
	NumGC      uint32 `json:"numgc,omitempty"`
}

type Ready struct {
	Ready string `json:"ready,omitempty"`
}

type Version struct {
	Version string     `json:"version,omitempty"`
	ReleaseDate string `json:"release-date,omitempty"`
}

var quotes []Quote
var authors []Author
var db *gorm.DB
var err error
var conf Configuration

// init manages initalization logic
// Uses envionment variable and configuration files.
// Gathers application run time settings
func init() {

	conf.Environment = os.Getenv("REST_API_ENV")
	if (conf.Environment == "") {
		fmt.Println("REST_API_ENV not defined, using default development environment settings.")
		conf.Environment = "development"
	}
	env := []string{}
	env = append(env, "config/config.", conf.Environment, ".json")
	err := gonfig.GetConf(strings.Join(env, ""), &conf)
	if (err != nil) {
		fmt.Sprintf("Environment %s file not found.", strings.Join(env, ""))
	}
}

// GetAuthors looks up all of the authors.
// GET /authors
// Populates authors slice with all of the author records in the database and returns JSON formatted listing.
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors = []Author{}
	db.Find(&authors)

	respondWithJSON(w, http.StatusOK, authors)
}

// GetAuthor looks up a specific author by ID.
// GET /author
// Looks up a author in the database by ID and returns results JSON format.
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author Author
	db.First(&author, params["id"])

	respondWithJSON(w, http.StatusOK, author)
}

// CreateAuthor creates a new author.
// POST /author
// Returns newly created author ID.
func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	db.Create(&author)

	respondWithJSON(w, http.StatusCreated, author)
	return
}

// GetQuotes looks up all of the quotes.
// GET /quotes
// Returns all of the quotes in the JSON format.
func GetQuotes(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, quotes)
}

// GetQuote looks up a specific quote by ID.
// GET /quote/{id}
// Returns a quote in the JSON format provided the target ID is valid.
func GetQuote(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//quoteID, err := strconv.Atoi(params["id"])
	//if (err != nil) {
	//	respondWithError(w, http.StatusBadRequest, "Invalid quote ID")
	//	return
	//}
	//
	//for _, item := range quotes {
	//	if item.ID == quoteID {
	//		respondWithJSON(w, http.StatusOK, item)
	//		return
	//	}
	//}
	//
	//// quoteID not found
	//respondWithError(w, http.StatusNotFound, "Quote not found.")
	return
}

// CreateQuote creates a new quote. Validates that the author ID exists.
// POST /quote
// Returns newly created quote. Including the ID of quote and the details of the author.
func CreateQuote(w http.ResponseWriter, r *http.Request) {

	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)

	// Validate that the author ID exists
	//var author Author
	//db.First(&author, quote.AuthorID)
	//if (author.ID == 0) {
	//	respondWithError(w, http.StatusBadRequest, "Invalid author ID")
	//	return
	//}

	db.Create(&quote)

	respondWithJSON(w, http.StatusCreated, quote)
}

// DeleteQuote deletes a quote by quote ID.
// DELETE /quote/{id}
// Returns all quotes which will exclude the deleted quote made by the DELETE request.
func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//quoteID, err := strconv.Atoi(params["id"])
	//if (err != nil) {
	//	respondWithError(w, http.StatusBadRequest, "Invalid quote ID")
	//	return
	//}
	//
	////for index, item := range quotes {
	////	if item.ID == quoteID {
	////		quotes = append(quotes[:index], quotes[:index+1]...)
	////		break
	////	}
	////}
	//// Add support for reporting quote not found
	//
	//respondWithJSON(w, http.StatusOK, quotes)
}

// GetHealth looks up the health of the application.
// GET /health
// Returns all of the health status of all the components of the application.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	var data Health

	runtime.ReadMemStats(&m)

	data.Refer = "https://golang.org/pkg/runtime/#MemStats"
	data.Alloc = m.Alloc / 1024
	data.TotalAlloc = m.TotalAlloc / 1024
	data.Sys = m.Sys / 1024
	data.NumGC = m.NumGC

	respondWithJSON(w, http.StatusOK, data)
}

// GetReady determines if the application is ready to process requests.
// GET /ready
// Returns the application state to determine if the application is ready to process requests.
func GetReady(w http.ResponseWriter, r *http.Request) {
	var data Ready
	data.Ready = "OK"
	respondWithJSON(w, http.StatusOK, data)
}

// GetVersion looks up the current version of the application.
// GET /version
// Returns the current version of the application.
func GetVersion(w http.ResponseWriter, r *http.Request) {
	var data Version
	configuration := Configuration{}

	env := []string{}
	env = append(env, "config/config.", configuration.Environment, ".json")
	err := gonfig.GetConf(strings.Join(env, ""), &configuration)
	if (err != nil) {
		respondWithError(w, http.StatusBadRequest, "Application Version details unknown!")
	}

	data.Version = configuration.Version
	data.ReleaseDate = configuration.ReleaseDate
	respondWithJSON(w, http.StatusOK, data)
}

// main function
// Starting point for application
func main() {

	// Postgres database
	fmt.Println("Starting DB connection...")
	db, err = gorm.Open(
		"postgres",
		"host=" + conf.DbHost + " " +
		"user=" + conf.DbUser + " " +
		"port=5432 " +
		"dbname=" + conf.DbName + " " +
		"sslmode=disable " +
		"password=" + conf.DbPassword)
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Quote{})
	db.AutoMigrate(&Author{})

	// API router
    // Consider use of .StrictSlash(true)
	router := mux.NewRouter()

	subRouterAuthors := router.PathPrefix("/authors").Subrouter()
	subRouterAuthor := router.PathPrefix("/author").Subrouter()
	//subRouterQuotes := router.PathPrefix("/quotes").Subrouter()
	subRouterQuote := router.PathPrefix("/quote").Subrouter()
	subRouterHealth := router.PathPrefix("/health").Subrouter()
	subRouterReady := router.PathPrefix("/ready").Subrouter()
	subRouterVersion := router.PathPrefix("/version").Subrouter()

	// GET /authors
	subRouterAuthors.HandleFunc("", GetAuthors).Methods("GET")
	subRouterAuthors.HandleFunc("/", GetAuthors).Methods("GET")

	// GET /author
	subRouterAuthor.HandleFunc("/{id}",  GetAuthor).Methods("GET")
	subRouterAuthor.HandleFunc("/{id}/", GetAuthor).Methods("GET")

	// POST /author
	subRouterAuthor.HandleFunc("", CreateAuthor).Methods("POST")
	subRouterAuthor.HandleFunc("/", CreateAuthor).Methods("POST")

	// DELETE /author
	//subRouterAuthor.HandleFunc("/{id}",  DeleteAuthor).Methods("DELETE")
	//subRouterAuthor.HandleFunc("/{id}/", DeleteAuthor).Methods("DELETE")

	// GET /quotes
	//subRouterQuotes.HandleFunc("", GetQuotes).Methods("GET")
	//subRouterQuotes.HandleFunc("/", GetQuotes).Methods("GET")

	// GET /quote
	//subRouterQuote.HandleFunc("/{id}",  GetQuote).Methods("GET")
	//subRouterQuote.HandleFunc("/{id}/", GetQuote).Methods("GET")

	// POST /quote
	subRouterQuote.HandleFunc("", CreateQuote).Methods("POST")
	subRouterQuote.HandleFunc("/", CreateQuote).Methods("POST")

	// DELETE /quote
	//subRouterQuote.HandleFunc("/{id}",  DeleteQuote).Methods("DELETE")
	//subRouterQuote.HandleFunc("/{id}/", DeleteQuote).Methods("DELETE")

	// GET /health
	subRouterHealth.HandleFunc("", GetHealth).Methods("GET")
	subRouterHealth.HandleFunc("/", GetHealth).Methods("GET")

	// GET /ready
	subRouterReady.HandleFunc("", GetReady).Methods("GET")
	subRouterReady.HandleFunc("/", GetReady).Methods("GET")

	// GET /version
	subRouterVersion.HandleFunc("", GetVersion).Methods("GET")
	subRouterVersion.HandleFunc("/", GetVersion).Methods("GET")

	if (conf.Port == 0) {
		fmt.Println("Application port setting not found")
		os.Exit(1)
	}
	port := []string{}
	port = append(port, ":", strconv.Itoa(conf.Port))

	fmt.Printf("Starting server on port %s\n", strings.Join(port, ""))
	log.Fatal(http.ListenAndServe(strings.Join(port, ""),
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(router))))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
