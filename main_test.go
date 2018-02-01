//
//
//
package main_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"net/http/httptest"
	"bytes"

	"github.com/deezone/rest-api-go/toolbox"
	"github.com/deezone/rest-api-go/application"
)

var a application.App

func TestMain(m *testing.M) {
	fmt.Println("Starting rest-api-go application in test mode...")

	// Set environment to "test"
	os.Setenv("REST_API_ENV", "test")

	a = application.App{}
	a.Initialize(
		toolbox.Conf.DbUser,
		toolbox.Conf.DbPassword,
		toolbox.Conf.DbName,
		toolbox.Conf.DbHost)

	code := m.Run()

	os.Exit(code)
}

//
func clearTables() {
	var quotes application.Quote
	var authors application.Author

	// Generally gorm simply updates the deleted_at field instead of deleting the record.
	// Delete record using ordinary deletion.
	a.DB.Unscoped().Delete(&quotes)
    a.DB.Raw("ALTER SEQUENCE quotes_id_seq RESTART WITH 1")

	a.DB.Unscoped().Delete(&authors)
	a.DB.Exec("ALTER SEQUENCE authors_id_seq RESTART WITH 1")

	return
}

//
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

//
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//
func TestEmptyQuotesTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/quotes", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Quote records not found." {
		t.Errorf("Expected the 'error' key of the response to be set to 'Quote records not found.'. Got '%s'", m["error"])
	}
}

//
func TestEmptyAuthorsTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/authors", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Author records not found." {
		t.Errorf("Expected the 'error' key of the response to be set to 'Author records not found.'. Got '%s'", m["error"])
	}
}

//
func TestGetNonExistentQuote(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/quote/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Quote ID: 1 not found." {
		t.Errorf("Expected the 'error' key of the response to be set to 'Quote ID: 1 not found.'. Got '%s'", m["error"])
	}
}

//
func TestGetNonExistentAuthor(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/author/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Author ID: 1 not found." {
		t.Errorf("Expected the 'error' key of the response to be set to 'Author ID: 1 not found.'. Got '%s'", m["error"])
	}
}

//
func TestCreateAuthor(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"first": 	   "First",
		"last":  	   "Last",
		"born":  	   "2000-01-02T00:00:00Z",
		"died":	 	   "2010-01-03T00:00:00Z",
		"description": "Test description.",
		"biolink": 	   "http://somesite.com"
	}`)

	req, _ := http.NewRequest("POST", "/author", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != "created" {
		t.Errorf("Expected the 'status' key of the response to be set to 'created'. Got '%s'", m["status"])
	}
	if m["id"] != "1" {
		t.Errorf("Expected the 'id' key of the response to be set to '1'. Got '%s'", m["id"])
	}
}

//
func TestCreateQuoteWithInvalidAuthor(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"quote":	"Test quote",
		"authorid":	666
	}`)

	req, _ := http.NewRequest("POST", "/quote", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Invalid author, authorid: 666 not found." {
		t.Errorf("Expected the 'error' key of the response to be set to 'Invalid author, authorid: 666 not found.'. Got '%s'", m["error"])
	}
}

//
func TestGetExistingAuthor(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"first": 	   "First",
		"last":  	   "Last",
		"born":  	   "2000-01-02T00:00:00Z",
		"died":	 	   "2010-01-03T00:00:00Z",
		"description": "Test description.",
		"biolink": 	   "http://somesite.com"
	}`)

	reqPOST, _ := http.NewRequest("POST", "/author", bytes.NewBuffer(payload))
	responsePOST := executeRequest(reqPOST)

	checkResponseCode(t, http.StatusCreated, responsePOST.Code)

	var m map[string]interface{}
	json.Unmarshal(responsePOST.Body.Bytes(), &m)

	if m["status"] != "created" {
		t.Errorf("Expected the 'status' key of the response to be set to 'created'. Got '%s'", m["status"])
	}
	if m["id"] != "1" {
		t.Errorf("Expected the 'id' key of the response to be set to '1'. Got '%s'", m["id"])
	}

	reqGET, _ := http.NewRequest("GET", "/author/1", nil)
	responseGET := executeRequest(reqGET)

	checkResponseCode(t, http.StatusOK, responseGET.Code)

	json.Unmarshal(responseGET.Body.Bytes(), &m)

	if m["first"] != "First" {
		t.Errorf("Expected the 'first' key of the response to be set to 'First'. Got '%s'", m["first"])
	}
	if m["last"] != "Last" {
		t.Errorf("Expected the 'last' key of the response to be set to 'Last'. Got '%s'", m["last"])
	}
	if m["born"] != "2000-01-01T19:00:00-05:00" {
		t.Errorf("Expected the 'born' key of the response to be set to '2000-01-01T19:00:00-05:00'. Got '%s'", m["born"])
	}
	if m["died"] != "2010-01-02T19:00:00-05:00" {
		t.Errorf("Expected the 'died' key of the response to be set to '2010-01-02T19:00:00-05:00'. Got '%s'", m["died"])
	}
	if m["description"] != "Test description." {
		t.Errorf("Expected the 'description' key of the response to be set to 'Test description.'. Got '%s'", m["description"])
	}
	if m["biolink"] != "http://somesite.com" {
		t.Errorf("Expected the 'bioloink' key of the response to be set to 'http://somesite.com'. Got '%s'", m["biolink"])
	}
}

//
func TestCreateQuoteWithValidAuthor(t *testing.T) {
	clearTables()

	payloadAuthor := []byte(`{
		"first": 	   "First",
		"last":  	   "Last",
		"born":  	   "2000-01-02T00:00:00Z",
		"died":	 	   "2010-01-03T00:00:00Z",
		"description": "Test description.",
		"biolink": 	   "http://somesite.com"
	}`)

	reqAuthor, _ := http.NewRequest("POST", "/author", bytes.NewBuffer(payloadAuthor))
	responseAuthor := executeRequest(reqAuthor)

	checkResponseCode(t, http.StatusCreated, responseAuthor.Code)

	var m map[string]interface{}
	json.Unmarshal(responseAuthor.Body.Bytes(), &m)

	if m["status"] != "created" {
		t.Errorf("Expected the 'status' key of the response to be set to 'created'. Got '%s'", m["status"])
	}
	if m["id"] != "1" {
		t.Errorf("Expected the 'id' key of the response to be set to '1'. Got '%s'", m["id"])
	}

	payloadQuote := []byte(`{
		"quote":	"Test quote",
		"authorid":	1
	}`)

	reqQuote, _ := http.NewRequest("POST", "/quote", bytes.NewBuffer(payloadQuote))
	responseQuote := executeRequest(reqQuote)

	checkResponseCode(t, http.StatusCreated, responseQuote.Code)

	json.Unmarshal(responseQuote.Body.Bytes(), &m)

	if m["status"] != "created" {
		t.Errorf("Expected the 'status' key of the response to be set to 'created'. Got '%s'", m["status"])
	}
	if m["id"] != "1" {
		t.Errorf("Expected the 'id' key of the response to be set to '1'. Got '%s'", m["id"])
	}
	if m["authorid"] != "1" {
		t.Errorf("Expected the 'authorid' key of the response to be set to '1'. Got '%s'", m["authorid"])
	}

	reqGET, _ := http.NewRequest("GET", "/author/1", nil)
	responseGET := executeRequest(reqGET)

	checkResponseCode(t, http.StatusOK, responseGET.Code)

	json.Unmarshal(responseGET.Body.Bytes(), &m)

	if m["first"] != "First" {
		t.Errorf("Expected the 'first' key of the response to be set to 'First'. Got '%s'", m["first"])
	}
	if m["last"] != "Last" {
		t.Errorf("Expected the 'last' key of the response to be set to 'Last'. Got '%s'", m["last"])
	}
	if m["born"] != "2000-01-01T19:00:00-05:00" {
		t.Errorf("Expected the 'born' key of the response to be set to '2000-01-01T19:00:00-05:00'. Got '%s'", m["born"])
	}
	if m["died"] != "2010-01-02T19:00:00-05:00" {
		t.Errorf("Expected the 'died' key of the response to be set to '2010-01-02T19:00:00-05:00'. Got '%s'", m["died"])
	}
	if m["description"] != "Test description." {
		t.Errorf("Expected the 'description' key of the response to be set to 'Test description.'. Got '%s'", m["description"])
	}
	if m["biolink"] != "http://somesite.com" {
		t.Errorf("Expected the 'bioloink' key of the response to be set to 'http://somesite.com'. Got '%s'", m["biolink"])
	}

	// @todo: test for quote values in response
}
