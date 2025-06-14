package sql

import (
	"database/sql"

	"braces.dev/errtrace"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func NewSQLite(dbPath string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errtrace.Wrap(err)
	}

	return &SQLite{db: db}, errtrace.Wrap(err)
}

func getDB(sqlite *SQLite) *sql.DB {
	if sqlite == nil || sqlite.db == nil {
		return nil
	}
	return sqlite.db
}
