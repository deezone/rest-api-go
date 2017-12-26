# rest-api-go
REST API using the Go programming language.

#### To Start
```
$ go run main.go
```

#### Endpoints
- GET `/quotes` -> All quotes in the quotes document (database)
- GET `/quote/{id}` -> Get a single quote
- POST `/quote/{id}` -> Create a new quote
- DELETE `/quote/{id}` -> Delete a quote

**POST Example**
- raw `JSON(application/json)`
```
{
    "id": 6,
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
- [Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)
- [Manage config in Golang to get variables from file and env variables](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)
- [Efficient String Concatenation in Go](http://herman.asia/efficient-string-concatenation-in-go)

##### Packages
- [time](https://golang.org/pkg/time/)
- [encoding/json](https://golang.org/pkg/encoding/json/)
- [strconv](https://golang.org/pkg/strconv/)
- [strings](https://golang.org/pkg/strings/)
- [os](https://golang.org/pkg/os/)

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
