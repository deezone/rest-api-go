[![GoDoc](https://godoc.org/github.com/DeeZone/rest-api-go?status.svg)](https://godoc.org/github.com/DeeZone/rest-api-go)

# rest-api-go
REST API using the Go programming language. A learning excercise to understand the use of the Go programming language
when building an API.

### To Start
```
$ go run main.go
```

### Hosted Version
- [rest-api-go on herokuapp](https://rest-api-go.herokuapp.com/)

### Endpoints
- GET `/quotes` -> All quotes in the quotes document (database)
- GET `/quote/{id}` -> Get a single quote
- POST `/quote/{id}` -> Create a new quote
- DELETE `/quote/{id}` -> Delete a quote

### Documentation
- [GoDoc Documentation](https://godoc.org/github.com/DeeZone/rest-api-go)

### Configuration

#### Environment Variables
- [Linux: Set Environment Variable](https://www.cyberciti.biz/faq/set-environment-variable-linux/)
```
export REST_API_ENV=development OR production
```

#### Configuration files
Add configuration files to:
- `config/config.development.json`
- `config/config.production.json`

*Sample*
```
{
  "Version": "v0.01",
  "ReleaseDate": "2018-01-06T18:00:00",
  "Port": 9001,
  "DbHost": "localhost",
  "DbUser": "dlee",
  "DbName": "rest-api-go",
  "DbPassword": ""
}
```

A combination of the environment variable and configuration file settings will be used by the application to determine
run time settings.

### Usage Example

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

### References
#### API Application in Go
- [Building a RESTful API with Golang](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
- [API foundations in Go](https://leanpub.com/api-foundations)
- [Get Programming with Go](https://www.manning.com/books/get-programming-with-go)
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
- [How I Built an API with Mux, Go, PostgreSQL, and GORM](https://dev.to/aspittel/how-i-built-an-api-with-mux-go-postgresql-and-gorm-5ah8)

#### Go Development
- [Your First Program](https://www.golang-book.com/books/intro/2)
- [Dep](https://github.com/golang/dep): dependency management tool for Go
- [Manage config in Golang to get variables from file and env variables](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)
- [Efficient String Concatenation in Go](http://herman.asia/efficient-string-concatenation-in-go)
- [Documenting Go Code With Godoc](https://www.goinggo.net/2013/06/documenting-go-code-with-godoc.html)
- [Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)
- [Getting Started with GORM](http://jinzhu.me/gorm/)
  - [GORM: A Simple Guide on CRUD](http://motion-express.com/blog/gorm:-a-simple-guide-on-crud)

#### Packages

- [encoding/json](https://golang.org/pkg/encoding/json/)
- [fmt](https://golang.org/pkg/fmt/)
- [log](https://golang.org/pkg/log/)
- [net/http](https://golang.org/pkg/net/http/)
  - [Status constants](https://golang.org/pkg/net/http/#pkg-constants)
- [os](https://golang.org/pkg/os/)
- [runtime](https://golang.org/pkg/runtime/)
- [strconv](https://golang.org/pkg/strconv/)
- [strings](https://golang.org/pkg/strings/)
- [time](https://golang.org/pkg/time/)

- [github.com/gorilla/mux](https://github.com/gorilla/mux)
  - [Gorilla Toolkit - Mux](http://www.gorillatoolkit.org/pkg/mux)
  - [Routing (using gorilla/mux)](https://gowebexamples.com/routes-using-gorilla-mux/)
- [github.com/gorilla/handlers](https://github.com/gorilla/handlers)
- [github.com/tkanos/gonfig](https://github.com/tkanos/gonfig)

#### Instrumentation / Performance
> _Instrumentation is a collective term for measuring instruments used for indicating, measuring and recording physical
quantities, and has its origins in the art and science of Scientific instrument-making._

- [10 Packages to Instrument Your Go Application](https://pocketgophers.com/10-to-instrument/)
- [Easy-metrics](https://github.com/admobi/easy-metrics): Ready to use, standalone Go metrics with interval snapshots
- [ocraft/health](https://github.com/gocraft/health)
> gocraft/health allows you to instrument your service for logging and metrics, and then send that instrumentation to
log files, StatsD, Bugsnag, or to be polled and aggregated via a JSON API.

> gocraft/health also ships with a New Relic-like aggregator (called healthd) that shows you your slowest endpoints,
top error producers, top throughput endpoints, and so on.

- [Log memory usage every n seconds in Go #golang](https://gist.github.com/thomas11/2909362)
- [Read Golang MemsStats the Hard Way](https://blog.xiaoba.me/2017/09/02/how-to-play-golang-with-gdb-and-python.html)

##### Go Application Hosting
- https://rest-api-go.herokuapp.com/
- [Getting Started on Heroku with Go](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)

##### Postgres
- [Postgres Documentation](https://www.postgresql.org/docs/)

##### Angular
- [Build A Full Stack Movie Database With Golang, Angular, And NoSQL](https://www.thepolyglotdeveloper.com/2017/02/build-a-full-stack-movie-database-with-golang-angular-and-nosql/)
- [Golang and Angular 2, Why not?](https://medium.com/@thanhngvpt/golang-and-angular-2-why-not-38a398b182c)
- [Create A Real Time Chat App With Golang, Angular 2, And Websockets](https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/)

##### Other
- [gobot.io](https://gobot.io/)
