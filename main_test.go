//
//
//
package main_test

import (
	"os"
	"fmt"
	"testing"

	"github.com/deezone/rest-api-go/app"
	"github.com/deezone/rest-api-go/toolbox"

	// "." in example article
	"github.com/deezone/rest-api-go"
)

var a main.App

func TestMain(m *testing.M) {
	fmt.Println("Starting rest-api-go application in test mode...")

	// @todo: confirm os.Getenv("REST_API_ENV") == 'test'

	a := main.App{}
	a.Initialize(
		os.Getenv(toolbox.Conf.DbUser),
		os.Getenv(toolbox.Conf.DbPassword),
		os.Getenv(toolbox.Conf.DbName),
		os.Getenv(toolbox.Conf.DbHost))

	a.Run(toolbox.Conf.DbHost)

	// ensureTableExists()

	code := m.Run()

	// clearTable()

	os.Exit(code)
}
