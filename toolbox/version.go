// The "ready" response functionality.
// A part of the  utility / toolbox methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"net/http"
	"strings"
)

type Version struct {
	Version string     `json:"version,omitempty"`
	ReleaseDate string `json:"release-date,omitempty"`
}

// init - one time initialization logic
func init() {
	fmt.Println("- toolbox/version application package initialized")
}

// GetVersion looks up the current version of the application.
// GET /version
// Returns the current version of the application.
func GetVersion(w http.ResponseWriter, r *http.Request) {
	var data Version
	configuration := Configuration{}

	env := []string{}
	env = append(env, "config/config.", configuration.Environment, ".json")
	err := gonfig.GetConf(strings.Join(env, ""), &configuration)
	if (err != nil) {
		RespondWithError(w, http.StatusBadRequest, "Application Version details unknown!")
	}

	data.Version = configuration.Version
	data.ReleaseDate = configuration.ReleaseDate
	RespondWithJSON(w, http.StatusOK, data)
}
