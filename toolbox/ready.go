// The "ready" response functionality for requests to the /ready endpoint.
// A part of the  toolbox (utility) methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"net/http"
)

type Ready struct {
	DbReady  string `json:"dbready,omitempty"`
	AppReady string `json:"appready,omitempty"`
}

// init - one time initialization logic
func init() {
	fmt.Println("- toolbox/ready application package initialized")
}

// GetReady determines if the application is ready to process requests.
// GET /ready
// Returns the application state to determine if the application is ready to process requests.
func GetReady(w http.ResponseWriter, r *http.Request) {
	var data Ready

	// Database state
	// Send a ping to make sure the database connection is alive.
	err = Db.DB().Ping()
	if err == nil {
		data.DbReady = "OK"
	} else {
		data.DbReady = "Error"
	}

	// Application state
	data.AppReady = "OK"

	RespondWithJSON(w, http.StatusOK, data)
}
