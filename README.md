[![GoDoc](https://godoc.org/github.com/DeeZone/rest-api-go?status.svg)](https://godoc.org/github.com/DeeZone/rest-api-go)

# rest-api-go
REST API using the Go programming language. A learning excercise to understand the use of the Go programming language
when building an API.

### To Start
```
$ go run main.go
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
