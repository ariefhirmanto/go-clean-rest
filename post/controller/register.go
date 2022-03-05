package controller

import (
	"project-go/post"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc post.Usecase) {
	h := NewPostHandlers(uc)

	postEndpoints := router.Group("/api/v1/post")

	{
		postEndpoints.POST("/", h.CreatePost)
		postEndpoints.GET("/", h.GetAllPost)
		postEndpoints.GET("/:id", h.GetPostByID)
		postEndpoints.GET("/title/:title", h.GetPostByTitle)
		postEndpoints.GET("/slug/:slug", h.GetPostBySlug)
		postEndpoints.DELETE("/:id", h.Delete)
		postEndpoints.PUT("/", h.UpdatePost)
	}
}
