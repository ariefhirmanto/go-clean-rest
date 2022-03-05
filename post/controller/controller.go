package controller

import (
	"net/http"
	"project-go/exception"
	"project-go/post"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type postHandlers struct {
	// cfg    *config.Config
	postUC post.Usecase
	// logger logger.Logger
}

type PostResponse struct {
	Message string `json:"message"`
}

// NewNewsHandlers News handlers constructor
func NewPostHandlers(postUC post.Usecase) *postHandlers {
	return &postHandlers{postUC: postUC}
}

func (h *postHandlers) CreatePost(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandler.Create")
	defer span.Finish()

	var input post.InputPostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	err := h.postUC.CreatePost(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, PostResponse{
		Message: "Success added post!",
	})
}

func (h *postHandlers) GetAllPost(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandlers.GetAllPost")
	defer span.Finish()

	post, err := h.postUC.FindAllPost(ctx)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandlers) GetPostByID(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandlers.GetPostByID")
	defer span.Finish()

	var input post.InputPostID
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	post, err := h.postUC.FindByID(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandlers) GetPostByTitle(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandlers.GetPostByTitle")
	defer span.Finish()

	var input post.InputPostTitle
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	post, err := h.postUC.FindByTitle(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandlers) GetPostBySlug(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandlers.GetPostBySlug")
	defer span.Finish()

	var input post.InputPostSlug
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	post, err := h.postUC.FindBySlug(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandlers) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandlers.DeletePost")
	defer span.Finish()

	var input post.InputPostID
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	err = h.postUC.DeletePost(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}

func (h *postHandlers) UpdatePost(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "postHandler.UpdatePost")
	defer span.Finish()

	var input post.InputUpdatePostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	err := h.postUC.UpdatePost(ctx, input)
	if err != nil {
		c.JSON(exception.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, PostResponse{
		Message: "Post successfully updated!",
	})
}
