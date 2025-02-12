package models

import (
	"errors"
	"strings"
	"time"
)

// Publication represents an user's publication
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare prepares publication data
func (publication *Publication) Prepare() error {
	if erro := publication.validation(); erro != nil {
		return erro
	}

	publication.format()
	return nil
}

// Validate validates publication data
func (publication *Publication) validation() error {
	if publication.Title == "" {
		return errors.New("title is required")
	}

	if publication.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

// format trims in publication data
func (publication *Publication) format() error {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)

	return nil
}
