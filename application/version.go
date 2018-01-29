// The "version" response functionality for requests to the /version endpoint.
// A part of the  toolbox (utility) methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package application

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"net/http"
	"strings"

	"github.com/deezone/rest-api-go/toolbox"
)

type Version struct {
	Version string     `json:"version,omitempty"`
	ReleaseDate string `json:"release-date,omitempty"`
}

// init - one time initialization logic
func init() {
	fmt.Println("- application/version rest-api-go package initialized")
}

// GetVersion looks up the current version of the application.
// GET /version
// Returns the current version of the application.
func GetVersion(w http.ResponseWriter, r *http.Request) {
	var data Version
	configuration := toolbox.Configuration{}

	env := []string{}
	env = append(env, "config/config.", configuration.Environment, ".json")
	err := gonfig.GetConf(strings.Join(env, ""), &configuration)
	if (err != nil) {
		toolbox.RespondWithError(w, http.StatusBadRequest, "Application Version details unknown!")
	}

	data.Version = configuration.Version
	data.ReleaseDate = configuration.ReleaseDate
	toolbox.RespondWithJSON(w, http.StatusOK, data)
}
