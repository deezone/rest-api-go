[![GoDoc](https://godoc.org/github.com/DeeZone/rest-api-go?status.svg)](https://godoc.org/github.com/DeeZone/rest-api-go)

# rest-api-go
REST API using the Go programming language. A learning excercise to understand the use of the Go programming language
when building an API.

### To Start
- ensure postGres database is running
```
pg-start
```

- create application databases
```
$ createdb rest-api-go
$ createdb test-rest-api-go
$ createdb dev-rest-api-go
```
NOTE: see [Installing Postgres via Brew](https://gist.github.com/sgnl/609557ebacd3378f3b72) for details on local setup.

- start application
```
$ go run main.go
- toolbox/config rest-api-go package initialized
REST_API_ENV not defined, using default development environment settings.
- toolbox/toolbox rest-api-go package initialized
- application/author rest-api-go package initialized
- application/authors rest-api-go package initialized
- application/health rest-api-go package initialized
- application/quote rest-api-go package initialized
- application/quotes rest-api-go package initialized
- application/ready rest-api-go package initialized
- application/version rest-api-go package initialized
Starting rest-api-go application...
Starting DB connection...
- application/application rest-api-go initializeRoutes() initialized
Starting server on port :9001
```

### Endpoints
- `GET /` -> Application details generate by Kong

- `GET /quotes`
- `GET /quote/{id}`
- `POST /quote`
```
{
	"quote": "Test quote.",
	"authorid": 1
}
```
- `DELETE /quote/{id}`

- `GET /quthors`
- `GET /author/{id}`
- `POST /author`
```
{
	"first": "First",
	"last": "Last",
	"born": "2000-01-02T00:00:00Z",
	"died": "2010-01-03T00:00:00Z",
	"description": "Test description.",
	"biolink": "http://somesite.com"
}
```
- `DELETE /quote{id}`

- `GET /status`
- `GET /health`
- `GET /verion`

### Configuration
- [Environment Variables](docs/configuration.md)
- [Configuration files](docs/configuration.md)

### References
- [Packages](docs/references.md)
- [API Application in Go](docs/references.md)
- [Go Development](docs/references.md)
- [Dockerize](docs/references.md)
- [Instrumentation/Performance](docs/references.md)
- [Go Application Hosting](docs/references.md)
- [Postgres](docs/references.md)
- [Angular](docs/references.md)

### Usage Examples
**Status**:
- [GET `/health/`](/docs/usage.md)
- [GET `/ready/`](/docs/usage.md)
- [GET `/version/`](/docs/usage.md)

**Core**:
- [GET `/`](/docs/usage.md)
- [GET `/quote/{id}`](/docs/usage.md)
- [GET `/quotes/`](/docs/usage.md)
- [GET `/author/{id}`](/docs/usage.md)
- [GET `/authors/`](/docs/usage.md)
- [POST `/quote/`](/docs/usage.md)
- [POST `/author/`](/docs/usage.md)
- [DELETE `/quote/{id}`](/docs/usage.md)
- [DELETE `/author/{id}`](/docs/usage.md)

### Hosted Version (pending)
- [rest-api-go on herokuapp](https://rest-api-go.herokuapp.com/)

### Documentation
- [GoDoc Documentation](https://godoc.org/github.com/DeeZone/rest-api-go)

### Test Coverage
```
$ go test -v
- toolbox/config rest-api-go package initialized
- toolbox/toolbox rest-api-go package initialized
- application/author rest-api-go package initialized
- application/authors rest-api-go package initialized
- application/health rest-api-go package initialized
- application/quote rest-api-go package initialized
- application/quotes rest-api-go package initialized
- application/ready rest-api-go package initialized
- application/version rest-api-go package initialized
Starting rest-api-go application in test mode...
Starting DB connection...
- application/application rest-api-go initializeRoutes() initialized
=== RUN   TestEmptyQuotesTable
--- PASS: TestEmptyQuotesTable (0.00s)
=== RUN   TestEmptyAuthorsTable
--- PASS: TestEmptyAuthorsTable (0.00s)
=== RUN   TestGetNonExistentQuote
--- PASS: TestGetNonExistentQuote (0.00s)
=== RUN   TestGetNonExistentAuthor
--- PASS: TestGetNonExistentAuthor (0.00s)
=== RUN   TestCreateAuthor
--- PASS: TestCreateAuthor (0.00s)
=== RUN   TestCreateQuoteWithInvalidAuthor
--- PASS: TestCreateQuoteWithInvalidAuthor (0.00s)
=== RUN   TestGetExistingAuthor
--- PASS: TestGetExistingAuthor (0.00s)
=== RUN   TestCreateQuoteWithValidAuthor
--- PASS: TestCreateQuoteWithValidAuthor (0.00s)
PASS
ok  	github.com/deezone/rest-api-go	0.090s
```
