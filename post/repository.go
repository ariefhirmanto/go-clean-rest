package post

import (
	"context"
	"project-go/models"
)

type Repository interface {
	CreatePost(ctx context.Context, input models.Post) (err error)
	GetAllPost(ctx context.Context) (resp []models.Post, err error)
	FindByID(ctx context.Context, ID int64) (resp models.Post, err error)
	FindByTitle(ctx context.Context, title string) (resp models.Post, err error)
	FindBySlug(ctx context.Context, slug string) (resp models.Post, err error)
	DeletePost(ctx context.Context, ID int64) error
	UpdatePost(ctx context.Context, input models.Post) (err error)
}
