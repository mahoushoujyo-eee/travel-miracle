package handler

import (
	"context"
	"travel/biz/param"
	"travel/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
)


var JudgeHandler = GenericHandler(func(ctx context.Context, c *app.RequestContext, request *param.JudgeRequest) (*int, error) {
		err := service.NewFeedbackService(ctx, c).StoreUserJudgement(request)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
