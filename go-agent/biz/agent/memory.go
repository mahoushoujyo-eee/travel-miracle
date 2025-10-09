package agent

import (
	"context"
	"encoding/json"
	"travel/biz/config"
	"travel/biz/model"

	"github.com/google/uuid"
)


func CreateConversation(ctx context.Context, userId int64) (string, error) {
	// 生成UUID作为conversationId
	conversationId := uuid.New().String()
	
	// 创建会话记录
	conversation := &model.Conversation{
		UserId:         userId,
		ConversationId: conversationId,
	}
	
	// 保存到数据库
	if err := config.DB.WithContext(ctx).Create(conversation).Error; err != nil {
		return "", err
	}
	
	return conversationId, nil
}

func UpdateConversationTitle(ctx context.Context, conversationId string, title string) error {
	// 更新会话标题
	if err := config.DB.WithContext(ctx).Model(&model.Conversation{}).
		Where("conversation_id = ?", conversationId).
		Update("title", title).Error; err != nil {
		return err
	}
	return nil
}

func GetConversationList(ctx context.Context, userId int64) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	if err := config.DB.WithContext(ctx).Model(&model.Conversation{}).
		Where("user_id = ?", userId).
		Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

func InsertMemory(ctx context.Context, conversationId string, eventType string, content string) error {
	// TODO: 处理类型转换
	memory := &model.ChatMemory{
		ConversationId: conversationId,
		Prompt:         content,
		Type:           eventType,
	}

	// 插入记录
	if err := config.DB.WithContext(ctx).Create(memory).Error; err != nil {
		return err
	}
	return nil
}

func InsertMemoryWithMetaData(ctx context.Context, conversationId string, eventType string, content string, metadata string) error {
	// TODO: 处理类型转换
	memory := &model.ChatMemory{
		ConversationId: conversationId,
		Prompt:         content,
		Type:           eventType,
		Metadata:       metadata,
	}

	// 插入记录
	if err := config.DB.WithContext(ctx).Create(memory).Error; err != nil {
		return err
	}
	return nil
}

func InsertMemoryWithImgs(ctx context.Context, conversationId string, eventType string, content string, imgUrls []string) error {
	// 将imgUrls序列化为JSON字符串
	imgUrlsJSON, err := json.Marshal(imgUrls)
	if err != nil {
		return err
	}
	
	memory := &model.ChatMemory{
		ConversationId: conversationId,
		Prompt:         content,
		Type:           eventType,
		Metadata:       string(imgUrlsJSON),
	}

	// 插入记录
	if err := config.DB.WithContext(ctx).Create(memory).Error; err != nil {
		return err
	}
	return nil
}

func InsertMemoryWithTool(ctx context.Context, conversationId string, eventType string, content string, tool string) error {
	// TODO: 处理类型转换
	memory := &model.ChatMemory{
		ConversationId: conversationId,
		Prompt:         content,
		Type:           eventType,
		Metadata:       tool,
	}

	// 插入记录
	if err := config.DB.WithContext(ctx).Create(memory).Error; err != nil {
		return err
	}
	return nil
}

func GetMemoryList(ctx context.Context, conversationId string) ([]*model.ChatMemory, error) {
	var memories []*model.ChatMemory
	if err := config.DB.WithContext(ctx).Model(&model.ChatMemory{}).
		Where("conversation_id = ?", conversationId).
		Find(&memories).Error; err != nil {
		return nil, err
	}
	return memories, nil
}