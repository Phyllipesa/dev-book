package repository

import (
	"api/src/models"
	"database/sql"
)

type publications struct {
	db *sql.DB
}

// NewRepositoryPublications creates a new repository publications instance
func NewRepositoryPublications(db *sql.DB) *publications {
	return &publications{db}
}

// Create inserts a publication in database
func (repository publications) Create(publication models.Publication) (uint64, error) {
	statement, erro := repository.db.Prepare(
		`insert into publications (title, content, author_id) values (?, ?, ?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(
		publication.Title, publication.Content, publication.AuthorID,
	)
	if erro != nil {
		return 0, erro
	}

	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil
}
