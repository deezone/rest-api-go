[![GoDoc](https://godoc.org/github.com/DeeZone/rest-api-go?status.svg)](https://godoc.org/github.com/DeeZone/rest-api-go)

# rest-api-go
REST API using the Go programming language. A learning excercise to understand the use of the Go programming language 
when building an API.

#### To Start
```
$ go run main.go
```

#### Hosted Version
- [rest-api-go on herokuapp](https://rest-api-go.herokuapp.com/)

#### Endpoints
- GET `/quotes` -> All quotes in the quotes document (database)
- GET `/quote/{id}` -> Get a single quote
- POST `/quote/{id}` -> Create a new quote
- DELETE `/quote/{id}` -> Delete a quote

#### Documentation
- [GoDoc Documentation](https://godoc.org/github.com/DeeZone/rest-api-go)

#### Configuration

##### Environment Variables
- [Linux: Set Environment Variable](https://www.cyberciti.biz/faq/set-environment-variable-linux/)
```
export REST_API_ENV=development OR production
```

##### Configuration files
Add configuration files to:
- `config/config.development.json`
- `config/config.production.json`

*Sample*
```
{
  "Port": 8001
}
```

A combination of the environment variable and configuration file settings will be used by the application to determine
run time settings.

#### Development Examples

- POST `/quote/{id}` -> Create a new quote

**_Example_**

`POST /quote/6`
- raw `JSON(application/json)`
```
{
    "quote": "Test quote.",
    "authour": {
        "id": 6,
        "first": "First",
        "last": "Last",
        "born": "2000-01-02T00:00:00Z",
        "died": "2010-01-03T00:00:00Z",
        "description": "Test description.",
        "biolink": "http://somesite.com"
    }
}
```

#### References
##### API Application in Go
- [Building a RESTful API with Golang](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
- [API foundations in Go](https://leanpub.com/api-foundations)

##### Go Development
- [Your First Program](https://www.golang-book.com/books/intro/2)
- [Dep](https://github.com/golang/dep): dependency management tool for Go
- [Manage config in Golang to get variables from file and env variables](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)
- [Efficient String Concatenation in Go](http://herman.asia/efficient-string-concatenation-in-go)
- [Documenting Go Code With Godoc](https://www.goinggo.net/2013/06/documenting-go-code-with-godoc.html)
- [Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)

##### Packages
- [time](https://golang.org/pkg/time/)
- [encoding/json](https://golang.org/pkg/encoding/json/)
- [strconv](https://golang.org/pkg/strconv/)
- [strings](https://golang.org/pkg/strings/)
- [os](https://golang.org/pkg/os/)
- [fmt](https://golang.org/pkg/fmt/)
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
  - [Gorilla Toolkit - Mux](http://www.gorillatoolkit.org/pkg/mux)
  - [Routing (using gorilla/mux)](https://gowebexamples.com/routes-using-gorilla-mux/)

##### Go Application Hosting
- https://rest-api-go.herokuapp.com/
- [Getting Started on Heroku with Go](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)
