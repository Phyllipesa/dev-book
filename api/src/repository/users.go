package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewRepositoryUsers creates a new repository users instance
func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a user in the database
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}
