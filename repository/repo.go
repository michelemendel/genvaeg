package repository

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/michelemendel/genvaeg/constants"
)

const dbName = "sqlite3"

type Repo struct {
	DB *sql.DB
}

func NewRepo() *Repo {
	dbFile := filepath.Join(os.Getenv(constants.ENV_DB_DIR_KEY), os.Getenv(constants.ENV_DB_NAME_KEY))

	db, err := sql.Open(dbName, dbFile)
	if err != nil {
		slog.Error(err.Error())
	}
	r := &Repo{db}
	r.DBConfig()
	return r
}

func (r *Repo) DBConfig() {
	r.DB.Exec("PRAGMA journal_mode = WAL")
	r.DB.Exec("PRAGMA foreign_keys = ON")
	// This doesn't work
	r.DB.Exec("PRAGMA busy_timeout = 5000")
}

func (r *Repo) Close() {
	r.DB.Close()
}

func (r *Repo) GetDB() *sql.DB {
	return r.DB
}
