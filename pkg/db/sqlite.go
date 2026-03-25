package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

/*
  pkg/db/sqlite.go
  This file is solely for connecting to SQLite and few other functions.
*/

func NewSQLiteConnection(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// OPTIONAL parameters
	// optimization for SQLite
	_, err = db.Exec(`
		PRAGMA journal_mode = WAL;
		PRAGMA synchronous = NORMAL;
		PRAGMA busy_timeout = 5000;
	`)
	if err != nil {
		return nil, err
	}

	// We limited the connections to save RAM on the 4GB PC
	db.SetMaxOpenConns(1)
	db.SetConnMaxIdleTime(time.Hour)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("database connection successful")
	return db, nil
}

// AutoMigrate creates the tables if they don't exist
func AutoMigrate(db *sql.DB) error {
	// SQL script to define our database structure
	query := `
	-- Table for customer information
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		full_name TEXT NOT NULL,
		phone TEXT,
		email TEXT
	);

	-- Table for vehicle information linked to a customer
	CREATE TABLE IF NOT EXISTS vehicles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER NOT NULL,
		make TEXT NOT NULL,
		model TEXT NOT NULL,
		year INTEGER,
		plate TEXT UNIQUE,
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	);
	`

	// Execute the SQL query
	_, err := db.Exec(query)

	fmt.Println("auto migration done successfully")
	return err
}
