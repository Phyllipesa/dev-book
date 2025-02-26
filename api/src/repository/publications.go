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

// FindPublication find publications that would appear in the user's feed
func (repository publications) FindPublications(userID uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
		select distinct p.*, u.nick from publications p 
		inner join users u on u.id = p.author_id 
		inner join followers s on p.author_id = s.user_id
		where u.id = ? or s.follower_id = ?`,
		userID,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication
		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// UpdatePublication update a publication
func (repository publications) Update(publicationID uint64, publication models.Publication) error {
	statement, erro := repository.db.Prepare(
		"update publications set title = ?, content = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publication.Title, publication.Content, publicationID); erro != nil {
		return erro
	}

	return nil
}

// DeletePublication delete a publication
func (repository publications) Delete(publicationID uint64) error {
	statement, erro := repository.db.Prepare(
		"delete from publications where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// FindByUserId find all publications of a specific user
func (repository publications) FindByUserId(userId uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
		select p.*, u.nick from publications p
		join users u on u.id = p.author_id
		where p.author_id = ?`,
		userId,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication
		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Like like an publication
func (repository publications) Like(publicationID uint64) error {
	statement, erro := repository.db.Prepare(
		"update publications set likes = likes + 1 where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// UnLike unlike an publication
func (repository publications) Unlike(publicationID uint64) error {
	statement, erro := repository.db.Prepare(
		`UPDATE publications SET likes = 
		 CASE 
		 	WHEN likes > 0 THEN likes -1
		 ELSE 0
		 END
		 WHERE id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}
