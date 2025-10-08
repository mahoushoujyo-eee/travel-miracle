package agent

import (
	"context"
	"log"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

// CreateUserMessage 创建用户消息，支持文本和图片
func CreateUserMessage(ctx context.Context, conversationId string, prompt string, imgUrls []string) (adk.Message, error) {
	if len(imgUrls) > 0 {
		multiContent := []schema.ChatMessagePart{
			{
				Type: schema.ChatMessagePartTypeText,
				Text: prompt,
			},
		}
		for _, url := range imgUrls {
			multiContent = append(multiContent, schema.ChatMessagePart{
				Type: schema.ChatMessagePartTypeImageURL,
				ImageURL: &schema.ChatMessageImageURL{
					URL:    url,
					Detail: schema.ImageURLDetailAuto,
				},
			})
		}

		message := &schema.Message{
			Role: schema.User,
			MultiContent: multiContent,
		}

		// 异步插入带图片的记忆
		go InsertMemoryWithImgs(ctx, conversationId, string(schema.User), prompt, imgUrls)
		
		return message, nil
	} else {
		message := &schema.Message{
			Role: schema.User,
			Content: prompt,
		}

		// 异步插入普通记忆
		go InsertMemory(ctx, conversationId, string(schema.User), prompt)
		
		return message, nil
	}
}


func GetHistoryMessageList(ctx context.Context, conversationId string, userId int64) ([]adk.Message, string, error) {
	var err error
	var messages []adk.Message
	log.Printf("GetHistoryMessageList, conversationId: %s, userId: %d", conversationId, userId)
	if conversationId == ""{
		conversationId, err = CreateConversation(ctx, userId)
		if err != nil {
			return nil, conversationId, err
		}
	}else{
		chatMemory, err := GetMemoryList(ctx, conversationId)
		if err != nil {
			return nil, conversationId, err
		}

		// TODO: Type Transformation
		for _, memory := range chatMemory {
			messages = append(messages, &schema.Message{
				Role: TransformMemoryRoleToMessage(memory.Type),
				Content: memory.Prompt,
			})
		}
	}

	return messages, conversationId, nil
}


func TransformMemoryRoleToMessage(role string) schema.RoleType {

	//TODO: tool info
	switch role {
	case "user":
		return schema.User
	case "assistant":
	case "stream-chat":
	case "stream-reasoning":
		return schema.Assistant
	// case "tool":
	// 	return schema.Tool
	default:
		return schema.User
	}
	return schema.User
}