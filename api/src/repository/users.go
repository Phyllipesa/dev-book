package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewRepositoryUsers creates a new repository users instance
func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a user in database
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}

// FindAllWithNameOrNick finds all users with a name or nick
func (repository users) FindAllWithNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, erro := repository.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID finds a user by id
func (repository users) FindUserById(ID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// Update updates a user in the database
func (repository users) UpdateUser(ID uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Delete deletes a user from the database
func (repository users) DeleteUser(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// FindByEmail finds a user by email
func (repository users) FindByEmail(email string) (models.User, error) {
	lines, erro := repository.db.Query(
		"select id, password from users where email = ?",
		email,
	)

	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

// FollowUser allows a user follow other
func (repository users) FollowUser(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// UnfollowUser allows a user unfollow other
func (repository users) UnfollowUser(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// FindFollowers find all followers of a user
func (repository users) FindFollowers(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at from users u
		inner join followers s on u.id = s.follower_id where s.user_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if erro = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

// FindFollowing find for all users you are following
func (repository users) FindFollowing(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(
		`select u.id, u.name, u.nick, u.email, u.created_at from users u
		 inner join followers s on u.id = s.user_id where s.follower_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var following []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		following = append(following, user)
	}

	return following, nil
}

// FindPassword get a user's password by userId
func (repository users) FindPassword(userID uint64) (string, error) {
	lines, erro := repository.db.Query(
		"select password from users where id = ?",
		userID,
	)
	if erro != nil {
		return "", erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

// UpdatePassword updates a user's password in the database
func (repository users) UpdatePassword(userID uint64, newPassword string) error {
	statement, erro := repository.db.Prepare(
		"update users set password = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(newPassword, userID); erro != nil {
		return erro
	}

	return nil
}
