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

// FindById finds publication by id
func (repository publications) FindById(publicationID uint64) (models.Publication, error) {
	lines, erro := repository.db.Query(`
		select p.*, u.nick from publications p
		inner join users u
		on u.id = p.author_id where p.id = ?`,
		publicationID,
	)
	if erro != nil {
		return models.Publication{}, erro
	}
	defer lines.Close()

	var publication models.Publication

	if lines.Next() {
		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}
