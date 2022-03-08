package post

import (
	"context"
	"project-go/models"
)

type CacheRepository interface {
	GetPostByID(ctx context.Context, key string) (resp models.Post, err error)
	SetPostByID(ctx context.Context, key string, seconds int, post models.Post) error
	DeletePostByID(ctx context.Context, key string) error
	GetPostByTitle(ctx context.Context, title string) (resp models.Post, err error)
	SetPostByTitle(ctx context.Context, key string, seconds int, post models.Post) error
	GetPostBySlug(ctx context.Context, key string) (resp models.Post, err error)
	SetPostBySlug(ctx context.Context, key string, seconds int, post models.Post) error
}
