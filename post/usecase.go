package post

import (
	"context"
	"project-go/models"
)

type Usecase interface {
	CreatePost(ctx context.Context, input InputPostRequest) (err error)
	FindAllPost(ctx context.Context) (resp []models.Post, err error)
	FindByID(ctx context.Context, input InputPostID) (resp models.Post, err error)
	FindByTitle(ctx context.Context, input InputPostTitle) (resp models.Post, err error)
	FindBySlug(ctx context.Context, input InputPostSlug) (resp models.Post, err error)
	DeletePost(ctx context.Context, input InputPostID) error
	UpdatePost(ctx context.Context, input InputUpdatePostRequest) (err error)
}
