package repository

import (
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewRepositoryUsers creates a new repository users instance
func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}
