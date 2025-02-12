package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the application
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createAt,omitempty"`
}

// Prepare prepares the user data
func (user *User) Prepare(stage string) error {
	if erro := user.validation(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

// Validate validates user data
func (user *User) validation(stage string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("invalid email")
	}

	if stage == "cadastro" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

// format trims in user data
func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "cadastro" {
		passwordWithHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}
