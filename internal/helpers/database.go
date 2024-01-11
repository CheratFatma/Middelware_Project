package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//Initialisation de la BDD
func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "tchipify.db")
	if err != nil {
		return nil, err
	}

	if err := InitTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

//initialisation des tables
func InitTables(db *sql.DB) error {
	// Users
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT UNIQUE,
        email TEXT UNIQUE
    )`)
	if err != nil {
		return err
	}

	//Songs
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY,
		title TEXT,
		author TEXT
	)`)
	if err != nil {
		return err
	}

	return nil
}
