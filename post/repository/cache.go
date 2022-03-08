package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"project-go/models"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type postCacheRepository struct {
	redisClient *redis.Client
}

const (
	basePrefix = "post"
)

func NewCacheRepository(redisClient *redis.Client) *postCacheRepository {
	return &postCacheRepository{redisClient: redisClient}
}

// func (c *postCacheRepository) GetAllPost(ctx context.Context) ([]*models.Post, error) {
// 	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.GetAllPost")
// 	defer span.Finish()

// 	keyCache := c.getKeyWithPrefix("all", "")
// 	postBytes, err := c.redisClient.Get(ctx, keyCache).Bytes()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "postCacheRepository.GetAllPost.redisClient.Get")
// 	}
// 	post := &models.Post{}
// 	if err = json.Unmarshal(postBytes, post); err != nil {
// 		return nil, errors.Wrap(err, "postCacheRepository.GetAllPost.json.Unmarshal")
// 	}

// 	return post, nil
// }

func (c *postCacheRepository) GetPostByID(ctx context.Context, key string) (models.Post, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.GetPostByID")
	defer span.Finish()

	keyCache := c.getKeyWithPrefix("id", key)
	postBytes, err := c.redisClient.Get(ctx, keyCache).Bytes()
	if err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostByID.redisClient.Get")
	}
	post := models.Post{}
	if err = json.Unmarshal(postBytes, &post); err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostByID.json.Unmarshal")
	}

	return post, nil
}

// Cache news item
func (c *postCacheRepository) SetPostByID(ctx context.Context, key string, seconds int, post models.Post) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.SetPostByID")
	defer span.Finish()

	postBytes, err := json.Marshal(post)
	if err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostByID.json.Marshal")
	}

	keyCache := c.getKeyWithPrefix("id", key)
	if err = c.redisClient.Set(ctx, keyCache, postBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostByID.redisClient.Set")
	}
	return nil
}

func (c *postCacheRepository) DeletePostByID(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.DeletePostByID")
	defer span.Finish()

	keyCache := c.getKeyWithPrefix("id", key)
	if err := c.redisClient.Del(ctx, keyCache).Err(); err != nil {
		return errors.Wrap(err, "postCacheRepository.DeletePostByID.redisClient.Del")
	}
	return nil
}

func (c *postCacheRepository) GetPostByTitle(ctx context.Context, key string) (models.Post, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.GetPostByTitle")
	defer span.Finish()

	keyCache := c.getKeyWithPrefix("title", key)
	postBytes, err := c.redisClient.Get(ctx, keyCache).Bytes()
	if err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostByTitle.redisClient.Get")
	}
	post := models.Post{}
	if err = json.Unmarshal(postBytes, &post); err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostByTitle.json.Unmarshal")
	}

	return post, nil
}

// Cache news item
func (c *postCacheRepository) SetPostByTitle(ctx context.Context, key string, seconds int, post models.Post) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.SetPostByTitle")
	defer span.Finish()

	postBytes, err := json.Marshal(post)
	if err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostByTitle.json.Marshal")
	}

	keyCache := c.getKeyWithPrefix("title", key)
	if err = c.redisClient.Set(ctx, keyCache, postBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostByTitle.redisClient.Set")
	}
	return nil
}

func (c *postCacheRepository) GetPostBySlug(ctx context.Context, key string) (models.Post, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.GetPostBySlug")
	defer span.Finish()

	keyCache := c.getKeyWithPrefix("slug", key)
	postBytes, err := c.redisClient.Get(ctx, keyCache).Bytes()
	if err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostBySlug.redisClient.Get")
	}
	post := models.Post{}
	if err = json.Unmarshal(postBytes, &post); err != nil {
		return models.Post{}, errors.Wrap(err, "postCacheRepository.GetPostBySlug.json.Unmarshal")
	}

	return post, nil
}

// Cache news item
func (c *postCacheRepository) SetPostBySlug(ctx context.Context, key string, seconds int, post models.Post) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postCacheRepository.SetPostBySlug")
	defer span.Finish()

	postBytes, err := json.Marshal(post)
	if err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostBySlug.json.Marshal")
	}

	keyCache := c.getKeyWithPrefix("slug", key)
	if err = c.redisClient.Set(ctx, keyCache, postBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "postCacheRepository.SetPostBySlug.redisClient.Set")
	}
	return nil
}

func (c *postCacheRepository) getKeyWithPrefix(category string, key string) string {
	return fmt.Sprintf("%s-%s: %s", basePrefix, category, key)
}
