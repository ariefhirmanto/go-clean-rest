package post

import (
	"errors"
	"fmt"
	"project-go/models"

	"github.com/gosimple/slug"
)

type InputPostRequest struct {
	Title    string `json:"title,omitempty" db:"title"`
	Content  string `json:"content,omitempty" db:"content"`
	ImageURL string `json:"image_url,omitempty" db:"image_url"`
	Category string `json:"category,omitempty" db:"category"`
}

type InputUpdatePostRequest struct {
	ID       int64  `json:"id,omitempty" db:"id"`
	Title    string `json:"title,omitempty" db:"title"`
	Content  string `json:"content,omitempty" db:"content"`
	ImageURL string `json:"image_url,omitempty" db:"image_url"`
	Category string `json:"category,omitempty" db:"category"`
}

type InputPostID struct {
	ID int64 `uri:"id" binding:"required"`
}

type InputPostTitle struct {
	Title string `uri:"title" binding:"required"`
}

type InputPostSlug struct {
	Slug string `uri:"slug" binding:"required"`
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

func (input InputUpdatePostRequest) ValidateInput() error {
	if input.ID == 0 {
		return errors.New("ID can't be empty")
	}

	if input.Title == "" {
		return errors.New("title can't be empty")
	}

	if input.Content == "" {
		return errors.New("content can't be empty")
	}

	return nil
}

func CreateInput(input InputPostRequest) models.Post {
	inputSlug := fmt.Sprintf("%s", input.Title)
	data := models.Post{
		Title:    input.Title,
		Slug:     slug.Make(inputSlug),
		Content:  input.Content,
		ImageURL: input.ImageURL,
		Category: input.Category,
	}

	return data
}

func CreateUpdateInput(input InputUpdatePostRequest) models.Post {
	inputSlug := fmt.Sprintf("%s", input.Title)
	data := models.Post{
		ID:       input.ID,
		Title:    input.Title,
		Slug:     slug.Make(inputSlug),
		Content:  input.Content,
		ImageURL: input.ImageURL,
		Category: input.Category,
	}

	return data
}
