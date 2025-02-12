package repository

import "database/sql"

type publications struct {
	db *sql.DB
}

// NewRepositoryPublications creates a new repository publications instance
func NewRepositoryPublications(db *sql.DB) *users {
	return &users{db}
}
