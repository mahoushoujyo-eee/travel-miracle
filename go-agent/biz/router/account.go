package router

import (
	"travel/biz/middleware"

	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterUser(r *route.RouterGroup) {

	publicRouter := r.Group("/public")
	{
		publicRouter.POST("/login", middleware.JwtMiddleware.LoginHandler)
		publicRouter.POST("/register", )
		publicRouter.POST("/reset/email", )
		publicRouter.POST("/reset/password", )
	}

	commonRouter := r.Group("/common", middleware.JwtMiddleware.MiddlewareFunc())
	{
		commonRouter.GET("/hello", )
	}

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/")
	}
}