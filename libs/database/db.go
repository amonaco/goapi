package database

import (
	"log"
	"time"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
)

var db *pg.DB

const migrationsPath = "/app/migrations"

// Start connects to the postgres database and performs migrations to the specified version
func Start(url string) {
	db = Connect(url)

	// Run database migrations
	Migrate()
}

// Start connects to the postgres database and performs migrations to the specified version
func Connect(url string) *pg.DB {
	// Parse configuration string
	options, err := pg.ParseURL(url)
	if err != nil {
		panic(err)
	}

	// Connection retry settings
	options.MaxRetries = 10
	options.MinRetryBackoff = 1 * time.Second
	options.MaxRetryBackoff = 10 * time.Second

	// Connect to database
	database := pg.Connect(options)
	if err := checkConn(database); err != nil {
		log.Fatal("Error connecting to the database:\n\t", err)
	}

	return database
}

func Migrate() {

	// Autodiscover migrations from directory
	c := migrations.NewCollection()
	if err := c.DiscoverSQLMigrations(migrationsPath); err != nil {
		log.Fatal("Error discovering migrations:\n\t", err)
	}

	// Init migration table
	_, _, err := c.Run(db, "init")
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if !ok {
			log.Fatal(err)
		}

		// Avoid possible duplicate table error
		if pgErr.Field('C') != "42P07" {
			log.Fatal(err)
		}
	}

	// Run migrations up
	oldVersion, newVersion, err := c.Run(db, "up")
	if err != nil {
		log.Fatal(err)
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("Migration version is: %d\n", oldVersion)
	}
}

// Get returns the database connection
func Get() *pg.DB {
	return db
}

// Check connection is up and query-able
func checkConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}
