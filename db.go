package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func dbInsertInformation(siteName string, status string, latency int64) {
	insertQuery := "INSERT INTO stats (siteName, status, latency) VALUES (?, ?, ?)"
	_, err := database.Exec(insertQuery, siteName, status, latency)
	if err != nil {
		loggerError("Something went wrong on writing values: " + err.Error())
	}
}

func dbClearOldInformation() {
	query := "DELETE FROM stats WHERE createdAt < datetime('now', '-7 days')"
	_, err := database.Exec(query)
	if err != nil {
		loggerError("An error occurred on clearing old information: " + err.Error())
	}
	// vacuum?
}

func dbMain() {
	var err error
	database, err = sql.Open("sqlite3", "./monitor.db")
	if err != nil {
		loggerError("Something went wrong on opening database: " + err.Error())
	}

	// Creating table
	query := `
	CREATE TABLE IF NOT EXISTS stats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		siteName TEXT,
		status TEXT,
		latency INTEGER,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = database.Exec(query)
	if err != nil {
		loggerError("Something went wrong on creating table: " + err.Error())
	}
}
