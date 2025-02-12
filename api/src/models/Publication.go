package models

import "time"

// Publication represents an user's publication
type Publication struct {
	ID         uint64    `json:"id.omitempty"`
	Title      string    `json:"title.omitempty"`
	Content    string    `json:"content.omitempty"`
	AuthorID   uint64    `json:"authorId.omitempty"`
	AuthorNick uint64    `json:"authorNick.omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt.omitempty"`
}
