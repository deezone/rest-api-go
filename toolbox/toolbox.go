// The "toolbox" functionality common to many parts of the application.
// A part of the  toolbox (utility) methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// init - one time initialization logic
func init() {
	fmt.Println("- toolbox/toolbox application package initialized")
}

// respondWithError generates JSON formatted response values for error messages
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON generates JSON formatted response values for non error messages
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
