// The application route definations.
// A part of the  routes/routes methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"github.com/deezone/rest-api-go/toolbox"
	"github.com/gorilla/mux"
)

//
//
//
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
	subRouterAuthors.HandleFunc("", GetAuthors).Methods("GET")
	subRouterAuthors.HandleFunc("/", GetAuthors).Methods("GET")

	// GET /author
	subRouterAuthor.HandleFunc("/{id}",  GetAuthor).Methods("GET")
	subRouterAuthor.HandleFunc("/{id}/", GetAuthor).Methods("GET")

	// POST /author
	subRouterAuthor.HandleFunc("", CreateAuthor).Methods("POST")
	subRouterAuthor.HandleFunc("/", CreateAuthor).Methods("POST")

	// DELETE /author
	subRouterAuthor.HandleFunc("/{id}",  DeleteAuthor).Methods("DELETE")
	subRouterAuthor.HandleFunc("/{id}/", DeleteAuthor).Methods("DELETE")

	// GET /quotes
	subRouterQuotes.HandleFunc("", GetQuotes).Methods("GET")
	subRouterQuotes.HandleFunc("/", GetQuotes).Methods("GET")

	// GET /quote
	subRouterQuote.HandleFunc("/{id}",  GetQuote).Methods("GET")
	subRouterQuote.HandleFunc("/{id}/", GetQuote).Methods("GET")

	// POST /quote
	subRouterQuote.HandleFunc("", CreateQuote).Methods("POST")
	subRouterQuote.HandleFunc("/", CreateQuote).Methods("POST")

	// DELETE /quote
	subRouterQuote.HandleFunc("/{id}",  DeleteQuote).Methods("DELETE")
	subRouterQuote.HandleFunc("/{id}/", DeleteQuote).Methods("DELETE")

	// GET /health
	subRouterHealth.HandleFunc("", toolbox.GetHealth).Methods("GET")
	subRouterHealth.HandleFunc("/", toolbox.GetHealth).Methods("GET")

	// GET /ready
	subRouterReady.HandleFunc("", toolbox.GetReady).Methods("GET")
	subRouterReady.HandleFunc("/", toolbox.GetReady).Methods("GET")

	// GET /version
	subRouterVersion.HandleFunc("", toolbox.GetVersion).Methods("GET")
	subRouterVersion.HandleFunc("/", toolbox.GetVersion).Methods("GET")
}
