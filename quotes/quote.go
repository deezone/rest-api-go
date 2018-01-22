// The "quote" GET, POST and DALETE response functionality for requests to the /quote endpoint.
// A part of the  quotes methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package quotes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/deezone/rest-api-go/toolbox"
)

// init - one time initialization logic
func init() {
	fmt.Println("- quotes/quote application package initialized")
}

// GetQuote looks up a specific quote by ID.
// GET /quote/{id}
// Returns a quote in the JSON format provided the target ID is valid.
func GetQuote(w http.ResponseWriter, r *http.Request) {
	var quote toolbox.Quote

	params := mux.Vars(r)
	quoteID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid quote ID")
		return
	}

	// Check that quote ID is valid
	if (toolbox.Db.First(&quote, quoteID).RecordNotFound()) {
		message := []string{}
		message = append(message, "Quote ID: ", strconv.Itoa(int(quoteID)), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}

	// Lookup quote author
	// @todo: ISSUE-16 - create parameter to trigger author lookup rather than the default
	authormin := toolbox.AuthorMin{}
	toolbox.Db.Raw("SELECT * FROM authors WHERE id = ? AND deleted_at IS NULL", quote.AuthorID).Scan(&authormin)
	quote.Author = authormin

	toolbox.RespondWithJSON(w, http.StatusOK, quote)
}

// CreateQuote creates a new quote. Validates that the author ID exists.
// POST /quote
// Returns the ID of new quote as a part of the "status" response message.
func CreateQuote(w http.ResponseWriter, r *http.Request) {

	message := []string{}
	var quote toolbox.Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)

	// Validate that the author ID exists
	var author toolbox.Author
	if (toolbox.Db.First(&author, quote.AuthorID).RecordNotFound()) {
		message = append(message, "Invalid author, authorid: ", strconv.Itoa(int(quote.AuthorID)), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}

	toolbox.Db.Create(&quote)
	message = append(message, "Quote ID: ", strconv.Itoa(int(quote.ID)), " created for authorID: ",
		strconv.Itoa(int(quote.AuthorID)), ".")
	toolbox.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": strings.Join(message, "")})
}

// DeleteQuote deletes a quote by quote ID.
// DELETE /quote/{id}
// Returns.
func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quoteID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid quote ID")
		return
	}

	message := []string{}
	var quote toolbox.Quote
	if (toolbox.Db.First(&quote, quoteID).RecordNotFound()) {
		message = append(message, "Quote ID: ", strconv.Itoa(quoteID), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}
	toolbox.Db.Delete(&quote)

	// @todo: remove author ID from quotes that reference the deleted author

	message = append(message, "Quote ID: ", strconv.Itoa(quoteID), " deleted.")
	toolbox.RespondWithJSON(w, http.StatusOK, map[string]string{"status": strings.Join(message, "")})
}
