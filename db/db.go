package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	if err = DB.Ping(); err != nil {
		panic("Failed to verify database connection: " + err.Error())
	}

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    username TEXT NOT NULL,
	    password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Failed to create tables: %v", err))
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER,               
	    FOREIGN KEY (user_id) REFERENCES users(id)                              
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Failed to create tables: %v", err))
	}
	createRegistrationsTable := `
CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY (event_id) REFERENCES events(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
)`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(fmt.Sprintf("Failed to create tables: %v", err))
	}
}
