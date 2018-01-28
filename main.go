// A learning exercise to understand the use of the Go programming language when building an API
// Started Christmas 2017
// Governed by the license that can be found in the LICENSE file
package main

import (
	"fmt"

	"github.com/deezone/rest-api-go/toolbox"
	"github.com/deezone/rest-api-go/application"
)

// main function
// Starting point for application
func main() {
	fmt.Println("Starting rest-api-go application...")

	a := application.App{}
	a.Initialize(
		toolbox.Conf.DbUser,
		toolbox.Conf.DbPassword,
		toolbox.Conf.DbName,
		toolbox.Conf.DbHost)

	a.Run(toolbox.Conf.Port)

	// https://github.com/jinzhu/gorm/issues/1427
	// Commented out as having this within the application/application.go package seems to result in the connection
	// closing
	defer a.DB.Close()
}
