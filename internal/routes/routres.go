package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/handlers"
	"github.com/zlAyl/my-go-blog/internal/middleware"
	"github.com/zlAyl/my-go-blog/internal/repositories"
	"gorm.io/gorm"
)

func RegisterAllRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())
	// 初始化所有仓库和处理器
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	userRoute := NewUserRouter(userHandler)
	userRoute.RegisterRoutes(r)

	postRepo := repositories.NewPostRepository(db)
	postHandler := handlers.NewPostHandler(postRepo)
	postRoute := NewPostRouter(postHandler)
	postRoute.RegisterRoutes(r)

	commentRepo := repositories.NewCommentRepository(db)
	commentHandler := handlers.NewCommentHandler(commentRepo)
	commentRoute := NewCommentRouter(commentHandler)
	commentRoute.RegisterRoutes(r)

	return r
}
