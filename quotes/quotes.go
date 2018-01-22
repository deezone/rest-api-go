// The "quotes" response functionality for requests to the /quotes endpoint.
// A part of the  quotes methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package quotes

import (
	"fmt"
	"net/http"

	"github.com/deezone/rest-api-go/toolbox"
)

// init - one time initialization logic
func init() {
	fmt.Println("- quotes/quotes application package initialized")
}

// GetQuotes looks up all of the quotes.
// GET /quotes
// Returns all of the quotes in JSON format.
func GetQuotes(w http.ResponseWriter, r *http.Request) {
	count := 0
	quotes := []toolbox.Quote{}
	toolbox.Db.Find(&quotes).Count(&count)
	if count == 0 {
		toolbox.RespondWithError(w, http.StatusOK, "Quote records not found.")
		return
	}

	// Lookup quote author
	// @todo: ISSUE-16 - create parameter to trigger author lookup rather than being the default response
	authormin := toolbox.AuthorMin{}
	for index, quote := range quotes {
		toolbox.Db.Raw("SELECT * FROM authors WHERE id = ? AND deleted_at IS NULL", quote.AuthorID).Scan(&authormin)
		quotes[index].Author = authormin
	}

	toolbox.RespondWithJSON(w, http.StatusOK, quotes)
}
