package DB

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(path string) error {
	if err := open(path); err != nil {
		return err
	}
	if err := initTable(); err != nil {
		return err
	}
	return nil
}

func open(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	return err
}

func initTable() error {
	query := "CREATE TABLE IF NOT EXISTS Friends (id INTEGER PRIMARY KEY, name TEXT, url TEXT, description TEXT, avatar TEXT)"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	query = "CREATE TABLE IF NOT EXISTS FriendApplications (id INTEGER PRIMARY KEY, name TEXT, url TEXT, description TEXT,email TEXT, avatar TEXT,changeToken TEXT,approveToken TEXT)"
	_, err = db.Exec(query)
	return err
}

func GetDB() *sql.DB {
	return db
}
