// The "toolbox" functionality common to many parts of the application.
// A part of the  utility / toolbox methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"os"
	"strings"

	"github.com/tkanos/gonfig"
)

// Configuration type, settings common to application
type Configuration struct {
	Port        int
	Version     string
	ReleaseDate string
	Environment string
	DbHost      string
	DbUser      string
	DbName      string
	DbPassword  string
}

var Conf Configuration

// init manages initalization logic
// Uses envionment variable and configuration files.
// Gathers application run time settings
func init() {
	fmt.Println("- toolbox/config application package initialized")

	Conf.Environment = os.Getenv("REST_API_ENV")
	if (Conf.Environment == "") {
		fmt.Println("REST_API_ENV not defined, using default development environment settings.")
		Conf.Environment = "development"
	}
	env := []string{}
	env = append(env, "config/config.", Conf.Environment, ".json")
	Err := gonfig.GetConf(strings.Join(env, ""), &Conf)
	if (Err != nil) {
		fmt.Sprintf("Environment %s file not found.", strings.Join(env, ""))
	}
}

