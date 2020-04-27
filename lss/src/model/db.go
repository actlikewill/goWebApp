package model

import "database/sql"

var db *sql.DB

// SetDatabase initializes the database connection object.
func SetDatabase(database *sql.DB) {
	db = database
}