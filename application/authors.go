// The "authors" response functionality for requests to the /authors endpoint.
// A part of the  authors methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package application

import (
	"fmt"
	"net/http"

	"github.com/deezone/rest-api-go/toolbox"
)

var authors []Author

// init - one time initialization logic
func init() {
	fmt.Println("- application/authors rest-api-go package initialized")
}

// GetAuthors looks up all of the authors.
// GET /authors
// Populates authors slice with all of the author records in the database and returns JSON formatted listing.
// @todo: exclude author information in quotes list with each author
func (a *App) GetAuthors(w http.ResponseWriter, r *http.Request) {
	count := 0
	authors = []Author{}
	a.DB.Find(&authors).Count(&count)
	if count == 0 {
		toolbox.RespondWithError(w, http.StatusOK, "Author records not found.")
		return
	}

	// Lookup author quotes
	// @todo: ISSUE-16 - create parameter to trigger author lookup rather than being the default response
	// @todo: ISSUE-17 - create parameter to include deleted quotes in response
	quotesmin := []QuoteMin{}
	for index, author := range authors {
		a.DB.Raw("SELECT * FROM quotes WHERE author_id = ? AND deleted_at IS NULL", author.ID).Scan(&quotesmin).Count(&count)
		if count > 0 {
			authors[index].Quotes = quotesmin
		}
	}

	toolbox.RespondWithJSON(w, http.StatusOK, authors)
}
