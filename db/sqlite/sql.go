package sqlite

import (
	"database/sql"

	"github.com/allyourbasepair/allbase/db"
)

// SQLite is a SQLite database backend for Allbase.
type SQLite struct {
	db *sql.DB
}

// Init creates the database if it doesn't exist.
func (sqlite *SQLite) Init() error {
	return nil
}

// Connect opens a connection to the database.
func (sqlite *SQLite) Connect() error {
	return nil
}

// Config returns the database configuration.
func (sqlite *SQLite) Config() (db.Config, error) {
	return db.Config{}, nil
}

// SetConfig sets the database configuration.
func (sqlite *SQLite) SetConfig(db.Config) error {
	return nil
}

// Query returns a Query struct containing an http.Response struct and accompanying metadata
func (sqlite *SQLite) Query(query *db.Query) error {
	return nil
}

// Close closes the database.
func (sqlite *SQLite) Close() error {
	return nil
}
