package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var scheme = `
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY,
  description TEXT,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP,
  done INTEGER DEFAULT 0
);
`

type Task struct {
	Id          int
	Description string
	Created_at  string
	Done        int
}

func InitDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "todo_store.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(scheme)
	if err != nil {
		return nil, err
	}

	return db, nil
}
