package pkg

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

// DB is a global handle for database connection (kept simple for this stage)
var DB *sql.DB

// ConnectDB opens and pings the Postgres database.
func ConnectDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=admin dbname=phonebook-clean sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	// ensure connection is alive
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping failed", err)
	}
	log.Println("connected to the database")
}
