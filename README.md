# rest-api-go
REST API using the Go programming language.

#### To Start
```
$ go run main.go
```

#### Endpoints
- GET `/quote (GET)` -> All quotes in the quotes document (database)
- GET `/quote/{id}` -> Get a single quote
- POST `/quote/{id}` -> Create a new quote
- DELETE `/quote/{id}` -> Delete a quote

#### References
- [Building a RESTful API with Golang](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
- [Your First Program](https://www.golang-book.com/books/intro/2)
- [Dep](https://github.com/golang/dep): dependency management tool for Go