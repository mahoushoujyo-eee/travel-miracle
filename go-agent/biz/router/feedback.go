package router

import (
	"travel/biz/handler"

	"github.com/cloudwego/hertz/pkg/route"
)
	
func RegisterFeedback(r *route.RouterGroup) {
	feedbackRouter := r.Group("/user")
	{
		feedbackRouter.POST("/judge", handler.JudgeHandler)
	}
}