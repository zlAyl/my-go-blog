package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/handlers"
	"github.com/zlAyl/my-go-blog/internal/middleware"
)

type PostRouter struct {
	postHandler *handlers.PostHandler
}

func NewPostRouter(handler *handlers.PostHandler) *PostRouter {
	return &PostRouter{postHandler: handler}
}

func (postRoute *PostRouter) RegisterRoutes(router *gin.Engine) {
	router.GET("/post/list", postRoute.postHandler.List)
	router.GET("/post/detail/:id", postRoute.postHandler.Detail)

	userRoute := router.Group("/post").Use(middleware.AuthJWTMiddleware())
	{
		userRoute.POST("/publish", postRoute.postHandler.Publish)
		userRoute.PATCH("/update/:id", postRoute.postHandler.Update)
		userRoute.DELETE("/del/:id", postRoute.postHandler.Delete)
	}
}
