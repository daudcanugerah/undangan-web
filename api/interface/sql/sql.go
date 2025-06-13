package sql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func getDB(sqlite *SQLite) *sql.DB {
	if sqlite == nil || sqlite.db == nil {
		return nil
	}
	return sqlite.db
}
