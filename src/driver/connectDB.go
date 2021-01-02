package driver

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ASSESR_DIR = "./assets/"
	DBNAME     = "appointment.db"
)

func ConnectDB() (isOK bool, db *sql.DB) {
	dbPath := ASSESR_DIR + DBNAME
	var isOpen bool
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	return isOpen, db
}
