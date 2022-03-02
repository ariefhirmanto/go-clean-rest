package post

import (
	"context"
	"project-go/models"
)

type CacheRepository interface {
	GetPostByID(ctx context.Context, ID int64) (resp models.Post, err error)
	GetPostByTitle(ctx context.Context, title string) (resp models.Post, err error)
	GetPostBySlug(ctx context.Context, ID int64) (resp models.Post, err error)
}
