package router

import (
	"travel/biz/handler"

	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterConversation(r *route.RouterGroup) {
	conversationRouter := r.Group("/user")
	{
		conversationRouter.POST("/list", handler.GetConversationListHandler)
		conversationRouter.POST("/messages", handler.GetMemoryListHandler)
	}					
}
