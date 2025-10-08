package router

import (
	"travel/biz/handler"
	// "travel/biz/middleware"

	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterChat(r *route.RouterGroup) {
	// chatRouter := r.Group("/user", middleware.JwtMiddleware.MiddlewareFunc())
	chatRouter := r.Group("/user")
	{
		chatRouter.GET("/conversations", )
		chatRouter.GET("/conversations/:id/messages", )
		chatRouter.POST("/conversation", )
		chatRouter.GET("/files", )
		chatRouter.POST("/files", handler.GetUploadUrlHandler)
		chatRouter.POST("/stream", handler.ChatHandler)
	}
}
