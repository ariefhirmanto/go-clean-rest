package models

import (
	"errors"
	"time"
)

type InputCommentRequest struct {
	Title   string `json:"title,omitempty" db:"title"`
	Content string `json:"content,omitempty" db:"content"`
	Post    PostResponse
}

type CommentResponse struct {
	ID        int64        `json:"product_id,omitempty" db:"id"`
	Title     string       `json:"title,omitempty" db:"title"`
	Content   string       `json:"content,omitempty" db:"content"`
	CreatedAt time.Time    `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at,omitempty" db:"updated_at"`
	Post      PostResponse `json:"post,omitempty"`
}

func (input InputCommentRequest) ValidateInput() error {
	if input.Title == "" {
		return errors.New("title can't be empty")
	}

	if input.Content == "" {
		return errors.New("content can't be empty")
	}

	return nil
}
