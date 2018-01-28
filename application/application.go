//
//
//
package application

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Quote type (more like an object), manages the details of a quote.
type QuoteMin struct {
	// gorm.Model
	ID        uint        `gorm:"primary;key";json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Quote     string      `json:"quote"`
	AuthorID  uint        `gorm:"index";json:"authorid"`
}

// Quote type (more like an object), manages the details of a quote.
type Quote struct {
	QuoteMin
	Author 	 AuthorMin	 `json:"author,omitempty"`
}

// Author type, referenced by core items: quotes, publications, etc.
type AuthorMin struct {
	// gorm.Model
	ID          uint       `gorm:"primary;key";json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time

	First       string     `json:"first,omitempty"`
	Last        string     `json:"last,omitempty"`
	Born        time.Time  `json:"born,omitempty"`
	Died        time.Time  `json:"died,omitempty"`
	Description string     `json:"description,omitempty"`
	BioLink     string     `json:"biolink,omitempty"`
}

// Author type, referenced by core items: quotes, publications, etc.
type Author struct {
	AuthorMin
	Quotes      []QuoteMin `json:"quotes,omitempty"` // One-To-Many
}

var Quotes []Quote
var Quotesmin []QuoteMin
var Authors []Author
var Authorsmin []AuthorMin
var err error

func (a *App) Initialize(dbUser, dbPassword, dbName, dbHost string) {
	var err error

	// Postgres database
	fmt.Println("Starting DB connection...")
	a.DB, err = gorm.Open(
		"postgres",
		"host=" + dbHost + " " +
			"user=" + dbUser + " " +
			"dbname=" + dbName + " " +
			"sslmode=disable " +
			"password=" + dbPassword)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	// go-sql-driver/mysql - Best practice
	// https://github.com/go-sql-driver/mysql/issues/461#issuecomment-227008369
	// Db.DB().SetConnMaxLifetime(time.Minute*5);
	// Db.DB().SetMaxIdleConns(0);
	// Db.DB().SetMaxOpenConns(5);

	// https://github.com/jinzhu/gorm/issues/1427
	// Commented out as having this within a package seems to result in the connection closing,
	// moved to main.go -> main()
	// defer a.DB.Close()

	a.DB.AutoMigrate(&Quote{})
	a.DB.AutoMigrate(&Author{})

	// Consider use of .StrictSlash(true)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(dbPort int) {
	port := []string{}
	var portStr = append(port, ":", strconv.Itoa(dbPort))
	fmt.Printf("Starting server on port %s\n", strings.Join(portStr, ""))

	log.Fatal(http.ListenAndServe(strings.Join(portStr, ""),
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(a.Router))))
}

// init - one time initialization logic
// Application route definitions
func (a *App) initializeRoutes() {
	fmt.Println("- routes/routes rest-api-go package initialized")

	subRouterAuthors := a.Router.PathPrefix("/authors").Subrouter()
	subRouterAuthor := a.Router.PathPrefix("/author").Subrouter()
	subRouterQuotes := a.Router.PathPrefix("/quotes").Subrouter()
	subRouterQuote := a.Router.PathPrefix("/quote").Subrouter()
	subRouterHealth := a.Router.PathPrefix("/health").Subrouter()
	subRouterReady := a.Router.PathPrefix("/ready").Subrouter()
	subRouterVersion := a.Router.PathPrefix("/version").Subrouter()

	// GET /authors
	subRouterAuthors.HandleFunc("", a.GetAuthors).Methods("GET")
	subRouterAuthors.HandleFunc("/", a.GetAuthors).Methods("GET")

	// GET /author
	subRouterAuthor.HandleFunc("/{id}",  a.GetAuthor).Methods("GET")
	subRouterAuthor.HandleFunc("/{id}/", a.GetAuthor).Methods("GET")

	// POST /author
	subRouterAuthor.HandleFunc("", a.CreateAuthor).Methods("POST")
	subRouterAuthor.HandleFunc("/", a.CreateAuthor).Methods("POST")

	// DELETE /author
	subRouterAuthor.HandleFunc("/{id}", a.DeleteAuthor).Methods("DELETE")
	subRouterAuthor.HandleFunc("/{id}/", a.DeleteAuthor).Methods("DELETE")

	// GET /quotes
	subRouterQuotes.HandleFunc("", a.GetQuotes).Methods("GET")
	subRouterQuotes.HandleFunc("/", a.GetQuotes).Methods("GET")

	// GET /quote
	subRouterQuote.HandleFunc("/{id}",  a.GetQuote).Methods("GET")
	subRouterQuote.HandleFunc("/{id}/", a.GetQuote).Methods("GET")

	// POST /quote
	subRouterQuote.HandleFunc("", a.CreateQuote).Methods("POST")
	subRouterQuote.HandleFunc("/", a.CreateQuote).Methods("POST")

	// DELETE /quote
	subRouterQuote.HandleFunc("/{id}",  a.DeleteQuote).Methods("DELETE")
	subRouterQuote.HandleFunc("/{id}/", a.DeleteQuote).Methods("DELETE")

	// GET /health
	subRouterHealth.HandleFunc("", GetHealth).Methods("GET")
	subRouterHealth.HandleFunc("/", GetHealth).Methods("GET")

	// GET /ready
	subRouterReady.HandleFunc("", a.GetReady).Methods("GET")
	subRouterReady.HandleFunc("/", a.GetReady).Methods("GET")

	// GET /version
	subRouterVersion.HandleFunc("", GetVersion).Methods("GET")
	subRouterVersion.HandleFunc("/", GetVersion).Methods("GET")
}
