// The application route definations.
// A part of the  routes/routes methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package routes

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/deezone/rest-api-go/authors"
	"github.com/deezone/rest-api-go/quotes"
	"github.com/deezone/rest-api-go/toolbox"
)

// init - one time initialization logic
// Application route definitions
func init() {
	fmt.Println("- routes/routes application package initialized")

	// API router
	// Consider use of .StrictSlash(true)
	router := mux.NewRouter()

	subRouterAuthors := router.PathPrefix("/authors").Subrouter()
	subRouterAuthor := router.PathPrefix("/author").Subrouter()
	subRouterQuotes := router.PathPrefix("/quotes").Subrouter()
	subRouterQuote := router.PathPrefix("/quote").Subrouter()
	subRouterHealth := router.PathPrefix("/health").Subrouter()
	subRouterReady := router.PathPrefix("/ready").Subrouter()
	subRouterVersion := router.PathPrefix("/version").Subrouter()

	// GET /authors
	subRouterAuthors.HandleFunc("", authors.GetAuthors).Methods("GET")
	subRouterAuthors.HandleFunc("/", authors.GetAuthors).Methods("GET")

	// GET /author
	subRouterAuthor.HandleFunc("/{id}",  authors.GetAuthor).Methods("GET")
	subRouterAuthor.HandleFunc("/{id}/", authors.GetAuthor).Methods("GET")

	// POST /author
	subRouterAuthor.HandleFunc("", authors.CreateAuthor).Methods("POST")
	subRouterAuthor.HandleFunc("/", authors.CreateAuthor).Methods("POST")

	// DELETE /author
	subRouterAuthor.HandleFunc("/{id}", authors.DeleteAuthor).Methods("DELETE")
	subRouterAuthor.HandleFunc("/{id}/", authors.DeleteAuthor).Methods("DELETE")

	// GET /quotes
	subRouterQuotes.HandleFunc("", quotes.GetQuotes).Methods("GET")
	subRouterQuotes.HandleFunc("/", quotes.GetQuotes).Methods("GET")

	// GET /quote
	subRouterQuote.HandleFunc("/{id}",  quotes.GetQuote).Methods("GET")
	subRouterQuote.HandleFunc("/{id}/", quotes.GetQuote).Methods("GET")

	// POST /quote
	subRouterQuote.HandleFunc("", quotes.CreateQuote).Methods("POST")
	subRouterQuote.HandleFunc("/", quotes.CreateQuote).Methods("POST")

	// DELETE /quote
	subRouterQuote.HandleFunc("/{id}",  quotes.DeleteQuote).Methods("DELETE")
	subRouterQuote.HandleFunc("/{id}/", quotes.DeleteQuote).Methods("DELETE")

	// GET /health
	subRouterHealth.HandleFunc("", toolbox.GetHealth).Methods("GET")
	subRouterHealth.HandleFunc("/", toolbox.GetHealth).Methods("GET")

	// GET /ready
	subRouterReady.HandleFunc("", toolbox.GetReady).Methods("GET")
	subRouterReady.HandleFunc("/", toolbox.GetReady).Methods("GET")

	// GET /version
	subRouterVersion.HandleFunc("", toolbox.GetVersion).Methods("GET")
	subRouterVersion.HandleFunc("/", toolbox.GetVersion).Methods("GET")

	fmt.Printf("Starting server on port %s\n", strings.Join(toolbox.Conf.PortStr, ""))
	log.Fatal(http.ListenAndServe(strings.Join(toolbox.Conf.PortStr, ""),
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(router))))
}
