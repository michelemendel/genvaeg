package repository

import (
	"fmt"
	"log/slog"
)

func (r *Repo) runStatements(sqlStmts map[string]string) {
	for name, sqlStmt := range sqlStmts {
		_, err := r.DB.Exec(sqlStmt)
		if err != nil {
			slog.Error(fmt.Sprintf("Error in stmt %q:%s\n%s\n", err, name, sqlStmt))
		}
	}
}

// --------------------------------------------------------------------------------
// DDL
func (r *Repo) DropTables() {
	var sqlStmts = make(map[string]string)

	r.DB.Exec("PRAGMA foreign_keys = OFF")

	// Tables
	sqlStmts["drop_users"] = `DROP TABLE IF EXISTS users;`
	sqlStmts["drop_urls"] = `DROP TABLE IF EXISTS urls;`

	// Indexes
	sqlStmts["drop_index_urls_fullurl"] = `DROP INDEX IF EXISTS idx_urls_fullurl;`

	r.runStatements(sqlStmts)

	r.DB.Exec("PRAGMA foreign_keys = ON")
}

func (r *Repo) CreateTables() {
	var sqlStmts = make(map[string]string)

	sqlStmts["create_users"] = `
	CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		hashedpassword TEXT NOT NULL,
		created_at INTEGER DEFAULT CURRENT_TIMESTAMP
	); `

	sqlStmts["create_urls"] = `
	CREATE TABLE IF NOT EXISTS urls (
		shorturlpath TEXT PRIMARY KEY,
		fullurl TEXT NOT NULL,
    fk_user_uuid INTEGER,
		created_at INTEGER DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (fk_user_uuid) REFERENCES users(uuid)
	);`

	r.runStatements(sqlStmts)
}

func (r *Repo) CreateIndexes() {
	var sqlStmts = make(map[string]string)

	sqlStmts["create_index_urls_fullurl"] = `
	CREATE UNIQUE INDEX IF NOT EXISTS idx_urls_fullurl ON urls(fullurl);`

	r.runStatements(sqlStmts)
}
