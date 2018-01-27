// The "config" functionality common to many parts of the application.
// A part of the  toolbox (utility) methods for the rest-api-go application.
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

// init - one time initialization logic
// Uses environment variables and configuration files
// Gathers application run time settings to populate application Conf
func init() {
	fmt.Println("- toolbox/config rest-api-go package initialized")

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

	fmt.Sprintf("Environment %s loaded - Port: %s, Environment: %s, dbName: %s", strings.Join(env, ""), Conf.Port, Conf.Environment, Conf.DbName)

	if (Conf.Port == 0) {
		fmt.Println("Application port setting not found")
		os.Exit(1)
	}
}
