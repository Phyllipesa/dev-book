package models

import (
	"errors"
	"strings"
	"time"
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
func (user *User) Prepare() error {
	if erro := user.validation(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

// Validate validates the user data
func (user *User) validation() error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

// format trims the user data
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
