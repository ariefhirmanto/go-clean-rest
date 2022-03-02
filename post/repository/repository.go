package repository

import (
	"context"
	"database/sql"
	"log"
	"project-go/models"

	"github.com/opentracing/opentracing-go"
)

type postRepository struct {
	PostDB *sql.DB
}

func NewRepository(db *sql.DB) *postRepository {
	return &postRepository{db}
}

func (r *postRepository) CreatePost(ctx context.Context, input models.Post) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postRepository.AddPost")
	defer span.Finish()

	_, err = r.PostDB.ExecContext(ctx, addPostQuery,
		input.Title,
		input.Slug,
		input.Content,
		input.ImageURL,
		input.Category,
	)
	if err != nil {
		log.Println("[Post][CreatePost][Repository] Problem to querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (r *postRepository) GetAllPost(ctx context.Context) (resp []models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postRepository.GetAllPost")
	defer span.Finish()

	rows, err := r.PostDB.QueryContext(ctx, getAllPostQuery)
	if err != nil {
		log.Println("[Post][GetAllPost][Repository] Problem to querying to db, err: ", err.Error())
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.ImageURL,
			&post.Category); err != nil {
			return resp, err
		}

		resp = append(resp, post)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *postRepository) FindByID(ctx context.Context, ID int64) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postRepository.FindByID")
	defer span.Finish()

	err = r.PostDB.QueryRowContext(ctx, getPostQuery, ID).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Slug,
		&resp.Content,
		&resp.ImageURL,
		&resp.Category,
	)
	if err != nil {
		log.Println("[Post][FindByID][Repository] Problem to querying to db, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (r *postRepository) FindByTitle(ctx context.Context, title string) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postRepository.FindByTitle")
	defer span.Finish()

	err = r.PostDB.QueryRowContext(ctx, getPostByTitle, title).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Slug,
		&resp.Content,
		&resp.ImageURL,
		&resp.Category,
	)
	if err != nil {
		log.Println("[Post][FindByTitle][Storage] Problem to querying to db, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (r *postRepository) FindBySlug(ctx context.Context, slug string) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postRepository.FindBySlug")
	defer span.Finish()

	err = r.PostDB.QueryRowContext(ctx, getPostBySlug, slug).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Slug,
		&resp.Content,
		&resp.ImageURL,
		&resp.Category,
	)
	if err != nil {
		log.Println("[Post][FindBySlug][Repository] Problem to querying to db, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (r *postRepository) DeletePost(ctx context.Context, ID int64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PostRepository.Update")
	defer span.Finish()

	result, err := r.PostDB.ExecContext(ctx, deletePostByID, ID)
	if err != nil {
		log.Println("[Post][DeletePost][Storage] Problem to querying to db, err: ", err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("[Post][DeletePost][Storage] Problem checking rows affected after querying, err: ", err.Error())
		return err
	}

	if rowsAffected == 0 {
		log.Println("[Post][DeletePost][Storage] No rows, err: ", err.Error())
		return err
	}

	return nil
}

func (r *postRepository) UpdatePost(ctx context.Context, input models.Post) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PostRepository.Update")
	defer span.Finish()

	_, err = r.PostDB.ExecContext(
		ctx,
		updatePostQuery,
		&input.ID,
		&input.Title,
		&input.Slug,
		&input.Content,
		&input.ImageURL,
		&input.Category,
	)
	if err != nil {
		log.Println("[Post][UpdatePost][Storage] Problem to querying to db, err: ", err.Error())
		return err
	}

	return nil
}
