package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to the global DB variable
	if err != nil {
		panic("Could not connect to database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	var err error
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err = DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create users table.")
	}
	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event tables!")
	}

	createRegisterationTable := `
		CREATE TABLE IF NOT EXISTS registerations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			event_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(event_id) REFERENCES events(id)
		)
	`
	_, err = DB.Exec(createRegisterationTable)
	if err != nil {
		panic("Could not create registeration tables!")
	}
}
