package usecase

import (
	"context"
	"log"
	"project-go/models"
	"project-go/post"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type usecase struct {
	// cfg       *config.Config
	postRepo post.Repository
	// redisRepo news.RedisRepository
	// logger    logger.Logger
}

// Post Usecase constructor
func NewPostUsecase(postRepo post.Repository) *usecase {
	return &usecase{postRepo: postRepo}
}

func (u *usecase) CreatePost(ctx context.Context, input post.InputPostRequest) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.CreatePost")
	defer span.Finish()

	// user, err := utils.GetUserFromCtx(ctx)
	// if err != nil {
	// 	return nil, httpErrors.NewUnauthorizedError(errors.WithMessage(err, "newsUC.Create.GetUserFromCtx"))
	// }

	err = input.ValidateInput()
	if err != nil {
		log.Println("[Post][CreatePost][Usecase] Problem to querying to db, err: ", err.Error())
		return errors.WithMessage(err, "postUC.Create.ValidateInput")
	}

	data := post.CreateInput(input)
	err = u.postRepo.CreatePost(ctx, data)
	if err != nil {
		log.Println("[Post][CreatePost][Usecase] Problem to querying to db, err: ", err.Error())
		return errors.WithMessage(err, "postUC.Create.QueryingProblem")
	}

	return nil
}

func (u *usecase) FindAllPost(ctx context.Context) (resp []models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.FindAllPost")
	defer span.Finish()

	resp, err = u.postRepo.GetAllPost(ctx)
	if err != nil {
		log.Println("[Post][FindAllPost][Usecase] Problem to querying to db, err: ", err.Error())
		return resp, errors.WithMessage(err, "postUC.FindAllPost.QueryingProblem")
	}

	return resp, nil
}

func (u *usecase) FindByID(ctx context.Context, input post.InputPostID) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.FindByID")
	defer span.Finish()

	resp, err = u.postRepo.FindByID(ctx, input.ID)
	if err != nil {
		log.Println("[Post][FindByID][Usecase] Problem to querying to db, err: ", err.Error())
		return resp, errors.WithMessage(err, "postUC.FindByID.QueryingProblem")
	}

	return resp, nil
}

func (u *usecase) FindByTitle(ctx context.Context, input post.InputPostTitle) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.FindByTitle")
	defer span.Finish()

	resp, err = u.postRepo.FindByTitle(ctx, input.Title)
	if err != nil {
		log.Println("[Post][FindByTitle][Usecase] Problem to querying to db, err: ", err.Error())
		return resp, errors.WithMessage(err, "postUC.FindByTitle.QueryingProblem")
	}

	return resp, nil
}

func (u *usecase) FindBySlug(ctx context.Context, input post.InputPostSlug) (resp models.Post, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.FindBySlug")
	defer span.Finish()

	resp, err = u.postRepo.FindBySlug(ctx, input.Slug)
	if err != nil {
		log.Println("[Post][FindBySlug][Usecase] Problem to querying to db, err: ", err.Error())
		return resp, errors.WithMessage(err, "postUC.FindBySlug.QueryingProblem")
	}

	return resp, nil
}

func (u *usecase) DeletePost(ctx context.Context, input post.InputPostID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.DeletePost")
	defer span.Finish()

	err := u.postRepo.DeletePost(ctx, input.ID)
	if err != nil {
		log.Println("[Post][DeletePost][Usecase] Problem to querying to db, err: ", err.Error())
		return errors.WithMessage(err, "postUC.DeletePost.QueryingProblem")
	}

	return nil
}

func (u *usecase) UpdatePost(ctx context.Context, input post.InputUpdatePostRequest) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postUsecase.UpdatePost")
	defer span.Finish()

	err = input.ValidateInput()
	if err != nil {
		log.Println("[Post][UpdatePost][Usecase] Problem to querying to db, err: ", err.Error())
		return errors.WithMessage(err, "postUC.UpdatePost.ValidateInput")
	}

	data := post.CreateUpdateInput(input)
	err = u.postRepo.UpdatePost(ctx, data)
	if err != nil {
		log.Println("[Post][UpdatePost][Usecase] Problem to querying to db, err: ", err.Error())
		return errors.WithMessage(err, "postUC.UpdatePost.QueryingProblem")
	}

	return nil
}
