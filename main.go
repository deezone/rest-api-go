// A learning exercise to understand the use of the Go programming language when building an API
// Started Christmas 2017
// Governed by the license that can be found in the LICENSE file
package main

import (
	"fmt"
	"github.com/deezone/rest-api-go/toolbox"

	_ "github.com/deezone/rest-api-go/routes"
)

// main function
// Starting point for application
func main() {
	fmt.Println("Starting rest-api-go application...")

	// https://github.com/jinzhu/gorm/issues/1427
	// Commented out in toolbox/db.go as setting within a package seems to result in the connection closing prematurely
	// @todo ISSUE-22 : Research defer db.Close() within package
	defer toolbox.Db.Close()
}
