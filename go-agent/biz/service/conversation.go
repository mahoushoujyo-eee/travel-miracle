package service

import (
	"context"
	"travel/biz/config"
	"travel/biz/model"
	"travel/biz/param"

	"github.com/cloudwego/hertz/pkg/app"
)

type ConversationService struct{
	ctx context.Context
	c *app.RequestContext
}

func NewConversationService(ctx context.Context, c *app.RequestContext) *ConversationService {
	return &ConversationService{ctx: ctx, c: c}
}

func (s *ConversationService) GetConversationList(request *param.ChatRequest) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	if err := config.DB.WithContext(s.ctx).Model(&model.Conversation{}).
		Where("user_id = ?", request.UserId).
		Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

func (s *ConversationService) GetMemoryList(request *param.ChatRequest) ([]*model.ChatMemory, error) {
	var memories []*model.ChatMemory
	if err := config.DB.WithContext(s.ctx).Model(&model.ChatMemory{}).
		Where("conversation_id = ?", request.ConversationId).
		Find(&memories).Error; err != nil {
		return nil, err
	}
	return memories, nil
}

func (s *ConversationService) DeleteConversation(request *param.ChatRequest) error {
	if err := config.DB.WithContext(s.ctx).Model(&model.Conversation{}).
		Where("id = ?", request.ConversationId).
		Delete(&model.Conversation{}).Error; err != nil {
		return err
	}
	if err := config.DB.WithContext(s.ctx).Model(&model.ChatMemory{}).
		Where("conversation_id = ?", request.ConversationId).
		Delete(&model.ChatMemory{}).Error; err != nil {
		return err
	}
	return nil
}