// The "ready" response functionality.
// A part of the  utility / toolbox methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"net/http"
)

type Ready struct {
	Ready string `json:"ready,omitempty"`
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
	data.Ready = "OK"
	RespondWithJSON(w, http.StatusOK, data)
}
