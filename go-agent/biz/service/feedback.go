package service

import (
	"context"
	"travel/biz/config"
	"travel/biz/model"
	"travel/biz/param"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedbackService struct {
	ctx context.Context
	c *app.RequestContext
}

func NewFeedbackService(ctx context.Context, c *app.RequestContext) *FeedbackService {
	return &FeedbackService{
		ctx: ctx,
		c: c,
	}
}

func (s *FeedbackService) StoreUserJudgement(judgeRequest *param.JudgeRequest) error {
	err :=config.DB.WithContext(s.ctx).Create(&model.Feedback{
		Score:        judgeRequest.Score,
		Judgement:    judgeRequest.Judgement,
		ConversationID: judgeRequest.ConversationID,
	}).Error

	if err != nil {
		return err
	}
	return nil
}