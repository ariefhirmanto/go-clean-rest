package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"project-go/config"
	postHandler "project-go/post/controller"
	postRepository "project-go/post/repository"
	postUsecase "project-go/post/usecase"
)

type Server struct {
	httpServer  *http.Server
	cfg         *config.MainConfig
	db          *sql.DB
	redisClient *redis.Client
}

func NewServer(cfg *config.MainConfig, db *sql.DB, cache *redis.Client) *Server {
	return &Server{
		cfg:         cfg,
		db:          db,
		redisClient: cache,
	}
}

func (s *Server) Run() error {
	router := gin.Default()
	router.Use(CORSMiddleware())

	postRepo := postRepository.NewRepository(s.db)
	postCache := postRepository.NewCacheRepository(s.redisClient)

	postUC := postUsecase.NewPostUsecase(postRepo, postCache)
	postHandler.RegisterHTTPEndpoints(router, postUC)

	s.httpServer = &http.Server{
		Addr:           s.cfg.Server.Address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
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
