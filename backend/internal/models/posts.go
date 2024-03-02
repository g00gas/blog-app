package models

import "time"

type Post struct {
	PostID    int       `db:"post_id" json:"post_id,omitempty"`
	Title     string    `db:"title" json:"title,omitempty"`
	Content   string    `db:"content" json:"content"`
	Author    string    `db:"author" json:"author"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreatePostRequest struct {
	Title   string  `db:"title" json:"title,omitempty" binding:"required"`
	Content *string `db:"content" json:"content" binding:"required"`
	Author  string  `db:"author" json:"author,omitempty" binding:"required"`
}
