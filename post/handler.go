package post

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreatePost(c *gin.Context)
	GetAllPost(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetPostByTitle(c *gin.Context)
	GetPostBySlug(c *gin.Context)
	Delete(c *gin.Context)
	UpdatePost(c *gin.Context)
}
