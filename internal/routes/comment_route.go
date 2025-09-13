package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/handlers"
	"github.com/zlAyl/my-go-blog/internal/middleware"
)

type CommentRouter struct {
	commentHandler *handlers.CommentHandler
}

func NewCommentRouter(commentHandler *handlers.CommentHandler) *CommentRouter {
	return &CommentRouter{commentHandler: commentHandler}
}

func (commentRouter *CommentRouter) RegisterRoutes(router *gin.Engine) {
	commentRoute := router.Group("/comment").Use(middleware.AuthJWTMiddleware())
	{
		commentRoute.POST("/publish/:id", commentRouter.commentHandler.Publish)
		commentRoute.GET("/list/:id", commentRouter.commentHandler.Lists)
	}
}
