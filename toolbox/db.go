// The "db" (database) functionality common to many parts of the application.
// A part of the toolbox (utility) methods for the rest-api-go application.
// Governed by the license that can be found in the LICENSE file
package toolbox

import (
	"fmt"
	"github.com/soulcycle/milestones/db"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

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
var Db *gorm.DB
var err error

// init - one time initialization logic
func init() {
	fmt.Println("- toolbox/db application package initialized")

	// Postgres database
	fmt.Println("Starting DB connection...")
	Db, err = gorm.Open(
		"postgres",
		"host=" + Conf.DbHost + " " +
			"user=" + Conf.DbUser + " " +
			"dbname=" + Conf.DbName + " " +
			"sslmode=disable " +
			"password=" + Conf.DbPassword)
	if err != nil {
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
	// defer Db.Close()

	Db.AutoMigrate(&Quote{})
	Db.AutoMigrate(&Author{})
}
