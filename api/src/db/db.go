package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connection returns a connection with the database
func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnectionDB)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
