package models

import (
	"errors"
	"time"
)

type InputPostRequest struct {
	Title    string `json:"title,omitempty" db:"title"`
	Content  string `json:"content,omitempty" db:"content"`
	ImageURL string `json:"image_url,omitempty" db:"image_url"`
	Category string `json:"category,omitempty" db:"price"`
}

type PostResponse struct {
	ID        int64     `json:"product_id,omitempty" db:"id"`
	Title     string    `json:"title,omitempty" db:"title"`
	Content   string    `json:"content,omitempty" db:"content"`
	ImageURL  string    `json:"image_url,omitempty" db:"image_url"`
	Category  string    `json:"category,omitempty" db:"price"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (input InputPostRequest) ValidateInput() error {
	if input.Title == "" {
		return errors.New("title can't be empty")
	}

	if input.Content == "" {
		return errors.New("content can't be empty")
	}

	return nil
}
