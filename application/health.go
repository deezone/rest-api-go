// The "health" response functionality for requests to the /health endpoint.
// A part of the  toolbox (utility) methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package application

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/deezone/rest-api-go/toolbox"
)

type Health struct {
	Refer      string `json:"reference,omitempty"`
	Alloc      uint64 `json:"alloc,omitempty"`
	TotalAlloc uint64 `json:"total-alloc,omitempty"`
	Sys        uint64 `json:"sys,omitempty"`
	NumGC      uint32 `json:"numgc,omitempty"`
}

// init - one time initialization logic
func init() {
	fmt.Println("- application/health rest-api-go package initialized")
}

// GetHealth looks up the health of the application.
// GET /health
// Returns all of the health status of all the components of the application.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	var data Health

	runtime.ReadMemStats(&m)

	data.Refer = "https://golang.org/pkg/runtime/#MemStats"
	data.Alloc = m.Alloc / 1024
	data.TotalAlloc = m.TotalAlloc / 1024
	data.Sys = m.Sys / 1024
	data.NumGC = m.NumGC

	toolbox.RespondWithJSON(w, http.StatusOK, data)
}
