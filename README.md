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
- GET `/` -> Application details generate by Kong
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

### References
#### API Application in Go
- [Building a RESTful API with Golang](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
- [API foundations in Go](https://leanpub.com/api-foundations)
- [Get Programming with Go](https://www.manning.com/books/get-programming-with-go)
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
- [Building and Testing a REST API in GoLang using Gorilla Mux and MySQL](https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
- [How I Built an API with Mux, Go, PostgreSQL, and GORM](https://dev.to/aspittel/how-i-built-an-api-with-mux-go-postgresql-and-gorm-5ah8)

#### Go Development
- [Your First Program](https://www.golang-book.com/books/intro/2)
- [GO by Example](https://gobyexample.com/)
- [SyntaxDB - GO](https://syntaxdb.com/ref/go/)
- [Dep](https://github.com/golang/dep): dependency management tool for Go
- [Manage config in Golang to get variables from file and env variables](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)
- [Efficient String Concatenation in Go](http://herman.asia/efficient-string-concatenation-in-go)
- [Documenting Go Code With Godoc](https://www.goinggo.net/2013/06/documenting-go-code-with-godoc.html)
- [Godoc: documenting Go code](https://blog.golang.org/godoc-documenting-go-code)
- [Getting Started with GORM](http://jinzhu.me/gorm/)
  - [GORM: A Simple Guide on CRUD](http://motion-express.com/blog/gorm:-a-simple-guide-on-crud)
- [How to organize the go struct, in order to save memory.](https://medium.com/@felipedutratine/how-to-organize-the-go-struct-in-order-to-save-memory-c78afcf59ec2)

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
  - [Manage config in Golang to get variables from file and env variables](https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152)

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

### Usage Examples

##### GET `/`
**Response Code**: 200
```
{
    "version": "0.11.1",
    "plugins": {
        "enabled_in_cluster": [
            "cors"
        ],
        "available_on_server": {
            "response-transformer": true,
            "correlation-id": true,

... etc

}
```
##### GET `/quote/{id}`
**Response Code**: 200
```
{
    "ID": 2,
    "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
    "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
    "DeletedAt": null,
    "quote": "Test quote2.",
    "AuthorID": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
        "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
        "DeletedAt": null,
        "first": "First",
        "last": "Last",
        "born": "2000-01-01T19:00:00-05:00",
        "died": "2010-01-02T19:00:00-05:00",
        "description": "Test description.",
        "biolink": "http://somesite.com"
    }
}
```
**Resonse Code**: 400
```
{
    "error": "Quote ID: 22 not found."
}
```
##### GET `/quotes/`
**Response Code**: 200
```
[
       {
           "ID": 2,
           "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
           "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
           "DeletedAt": null,
           "quote": "Test quote2.",
           "AuthorID": 1,
           "author": {
               "ID": 1,
               "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
               "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
               "DeletedAt": null,
               "first": "First",
               "last": "Last",
               "born": "2000-01-01T19:00:00-05:00",
               "died": "2010-01-02T19:00:00-05:00",
               "description": "Test description.",
               "biolink": "http://somesite.com"
           }
       }
   ]

```
##### GET `/author/{id}`
**Response Code**: 200
```
{
    "ID": 1,
    "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
    "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
    "DeletedAt": null,
    "first": "First",
    "last": "Last",
    "born": "2000-01-01T19:00:00-05:00",
    "died": "2010-01-02T19:00:00-05:00",
    "description": "Test description.",
    "biolink": "http://somesite.com",
    "quotes": [
        {
            "ID": 1,
            "CreatedAt": "2018-01-20T22:36:10.788467-05:00",
            "UpdatedAt": "2018-01-20T22:36:10.788467-05:00",
            "DeletedAt": null,
            "quote": "Test quote1.",
            "AuthorID": 1
        },
        {
            "ID": 2,
            "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
            "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
            "DeletedAt": null,
            "quote": "Test quote2.",
            "AuthorID": 1
        }
    ]
}
```
**Response Code**: 400
```
{
    "error": "Author ID: 11 not found."
}
```
##### GET `/authors/`
**Response Code**: 200
```
[
    {
        "ID": 1,
        "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
        "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
        "DeletedAt": null,
        "first": "First",
        "last": "Last",
        "born": "2000-01-01T19:00:00-05:00",
        "died": "2010-01-02T19:00:00-05:00",
        "description": "Test description.",
        "biolink": "http://somesite.com",
        "quotes": [
            {
                "ID": 2,
                "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
                "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
                "DeletedAt": null,
                "quote": "Test quote2.",
                "AuthorID": 1
            }
        ]
    }
]
```

##### POST `/quote/`
**Response Code**: 201
**BODY**
```
{
	"quote": "Test quote3.",
	"authorid": 1
}
```
```
{
    "status": "Quote ID: 3 created for authorID: 1."
}
```
##### POST `/author/`
**Response Code**: 201
**BODY**
```
{
	"first": "First",
	"last": "Last2",
	"born": "2000-01-02T00:00:00Z",
	"died": "2010-01-03T00:00:00Z",
	"description": "Test description.",
	"biolink": "http://somesite.com"
}
```
```
{
    "status": "Author ID: 2 created."
}
```

##### DELETE `/quote/{id}`
**Response Code**: 200
```
{
    "status": "Quote ID: 1 deleted."
}
```
##### DELETE `/author/{id}`
**Response Code**: 200
```
{
    "status": "Author ID: 1 deleted."
}
```

##### GET `/health/`
```
{
    "reference": "https://golang.org/pkg/runtime/#MemStats",
    "alloc": 3495,
    "total-alloc": 4304,
    "sys": 9030,
    "numgc": 8
}
```
##### GET `/ready/`
```
{
    "ready": "OK"
}
```
##### GET `/version/`
```
{
    "version": "v0.01",
    "release-date": "2018-01-06T18:00:00"
}
```
