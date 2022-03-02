package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"project-go/config"
	_postHandler "project-go/post/controller"
	_postRepo "project-go/post/repository"
	_postUsecase "project-go/post/usecase"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConfig := config.Database

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := sql.Open(`mysql`, connection)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	// repository
	postRepository := _postRepo.NewRepository(db)

	// usecase
	postUC := _postUsecase.NewPostUsecase(postRepository)

	// handler / controller
	postHandler := _postHandler.NewPostHandlers(postUC)

	// CORS
	router := gin.Default()
	router.Use(CORSMiddleware())
	api := router.Group("api/v1")

	// Router
	api.POST("/post", postHandler.CreatePost)
	api.GET("/post", postHandler.GetAllPost)
	api.GET("/post/:id", postHandler.GetPostByID)
	api.GET("/post/title/:title", postHandler.GetPostByTitle)
	api.GET("/post/slug/:slug", postHandler.GetPostBySlug)
	api.DELETE("/post/:id", postHandler.Delete)
	api.PUT("/post", postHandler.UpdatePost)
	// fmt.Printf("%+v\n", db)
	router.Run(config.Server.Address)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
