// The "author" GET, POST and DELETE response functionality for requests to the /author endpoint.
// A part of the  authors methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package authors

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
	fmt.Println("- authors/author application package initialized")
}

// GetAuthor looks up a specific author by ID.
// GET /author
// Looks up a author in the database by ID and returns results JSON format.
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	var author toolbox.Author

	params := mux.Vars(r)
	authorID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	// Check that author ID is valid
	if (toolbox.Db.First(&author, authorID).RecordNotFound()) {
		message := []string{}
		message = append(message, "Author ID: ", strconv.Itoa(int(authorID)), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}

	// Lookup author quotes
	// @todo: ISSUE-16 - create parameter to trigger this lookup rather than being the default
	// @todo: ISSUE-17 - create parameter to include deleted quotes in response
	quotesmin := []toolbox.QuoteMin{}
	toolbox.Db.Raw("SELECT * FROM quotes WHERE author_id = ? AND deleted_at IS NULL", author.ID).Scan(&quotesmin)
	author.Quotes = quotesmin

	toolbox.RespondWithJSON(w, http.StatusOK, author)
}

// CreateAuthor creates a new author.
// POST /author
// Returns newly created author ID.
func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author toolbox.Author
	_ = json.NewDecoder(r.Body).Decode(&author)

	// Create new record
	if err := toolbox.Db.Create(&author).Error; err != nil {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Error creatng author record.")
		return
	}

	message := []string{}
	message = append(message, "Author ID: ", strconv.Itoa(int(author.ID)), " created.")
	toolbox.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": strings.Join(message, "")})
}

// Delete Author deletes an author by author ID.
// DELETE /author/{id}
// Returns a status message that includes the ID of the author record deleted.
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	authorID, err := strconv.Atoi(params["id"])
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	message := []string{}
	var author toolbox.Author
	if (toolbox.Db.First(&author, authorID).RecordNotFound()) {
		message = append(message, "Author ID: ", strconv.Itoa(authorID), " not found.")
		toolbox.RespondWithError(w, http.StatusBadRequest, strings.Join(message, ""))
		return
	}
	toolbox.Db.Delete(&author)

	// @todo: ISSUE-18 - delete quotes attributed to deleted author

	message = append(message, "Author ID: ", strconv.Itoa(authorID), " deleted.")
	toolbox.RespondWithJSON(w, http.StatusOK, map[string]string{"status": strings.Join(message, "")})
}
