// The "author" GET, POST and DELETE response functionality for requests to the /author endpoint.
// A part of the  authors methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package application

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
	fmt.Println("- application/author rest-api-go package initialized")
}

// GetAuthor looks up a specific author by ID.
// GET /author
// Looks up a author in the database by ID and returns results JSON format.
func (a *App) GetAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author

	params := mux.Vars(r)
	authorID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	// Check that author ID is valid
	if (a.DB.First(&author, authorID).RecordNotFound()) {
		message := []string{}
		message = append(message, "Author ID: ", strconv.Itoa(int(authorID)), " not found.")
		toolbox.RespondWithError(w, http.StatusNotFound, strings.Join(message, ""))
		return
	}

	// Lookup author quotes
	// @todo: ISSUE-16 - create parameter to trigger this lookup rather than being the default
	// @todo: ISSUE-17 - create parameter to include deleted quotes in response
	quotesmin := []QuoteMin{}
	a.DB.Raw("SELECT * FROM quotes WHERE author_id = ? AND deleted_at IS NULL", author.ID).Scan(&quotesmin)
	author.Quotes = quotesmin

	toolbox.RespondWithJSON(w, http.StatusOK, author)
}

// CreateAuthor creates a new author.
// POST /author
// Returns newly created author ID.
func (a *App) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author

	_ = json.NewDecoder(r.Body).Decode(&author)

	// Create new record
	if err := a.DB.Create(&author).Error; err != nil {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Error creating author record.")
		return
	}

	// @todo: check if author already exists by "first", "last", "born", "died" values
	// https://code.i-harness.com/en/q/3a6146
	// respond with "status" : "already exists", "id": existing ID, http.Conflict: 409
	// or StatusUnprocessableEntity: 422

	m := make(map[string]string)
	m["status"] = "created"
	m["id"] = strconv.Itoa(int(author.ID))
	toolbox.RespondWithJSON(w, http.StatusCreated, m)
}

// Delete Author deletes an author by author ID.
// DELETE /author/{id}
// Returns a status message that includes the ID of the author record deleted.
func (a *App) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	authorID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	message := []string{}
	var author Author
	if (a.DB.First(&author, authorID).RecordNotFound()) {
		message = append(message, "Author ID: ", strconv.Itoa(authorID), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}
	a.DB.Delete(&author)

	// @todo: ISSUE-18 - delete quotes attributed to deleted author

	message = append(message, "Author ID: ", strconv.Itoa(authorID), " deleted.")
	toolbox.RespondWithJSON(w, http.StatusOK, map[string]string{"status": strings.Join(message, "")})
}
