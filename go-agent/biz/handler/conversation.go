package handler

import (
	"context"
	"travel/biz/model"
	"travel/biz/param"
	"travel/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
)

var GetConversationListHandler = GenericHandler(func(ctx context.Context, c *app.RequestContext, request *param.ChatRequest) (*[]*model.Conversation, error) {
		result, err := service.NewConversationService(ctx, c).GetConversationList(request)
		if err != nil {
			return nil, err
		}
		return &result, nil
	})

var GetMemoryListHandler = GenericHandler(func(ctx context.Context, c *app.RequestContext, request *param.ChatRequest) (*[]*model.ChatMemory, error) {
		result, err := service.NewConversationService(ctx, c).GetMemoryList(request)
		if err != nil {
			return nil, err
		}
		return &result, nil
	})
