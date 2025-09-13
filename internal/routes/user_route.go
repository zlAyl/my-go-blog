package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/handlers"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
}

func NewUserRouter(handler *handlers.UserHandler) *UserRouter {
	return &UserRouter{userHandler: handler}
}

func (uRouter *UserRouter) RegisterRoutes(router *gin.Engine) {
	userRoute := router.Group("/user")
	{
		userRoute.POST("/register", uRouter.userHandler.Register)
		userRoute.POST("/login", uRouter.userHandler.Login)
	}
}
