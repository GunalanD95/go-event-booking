package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func createTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		price INTEGER NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	regTable := `
	CREATE TABLE IF NOT EXISTS registers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
		FOREIGN KEY (event_id) REFERENCES events(id)
	);
	`
	_, err = DB.Exec(regTable)
	if err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	return nil
}

func InitDB() error {
	var err error
	db, err := sql.Open("sqlite", "event.db") // Ensure "sqlite" driver is used.
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTables()
	if err != nil {
		return fmt.Errorf("could not create tables: %w", err)
	}

	fmt.Println("Tables created successfully!")
	return nil
}
