package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

// CreateUserMessage 创建用户消息，支持文本和图片
func CreateUserMessageAndStore(ctx context.Context, conversationId string, prompt string, imgUrls []string) (adk.Message, error) {
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
			Role:         schema.User,
			MultiContent: multiContent,
		}

		// 异步插入带图片的记忆
		go func() {
			imgUrlsStr, _ := json.Marshal(imgUrls)
			InsertMemoryWithMetaData(ctx, conversationId, string(schema.User), prompt, "img-text")
			InsertMemoryWithMetaData(ctx, conversationId, string(schema.User), string(imgUrlsStr), "img-urls")
		}()
		return message, nil
	} else {
		message := &schema.Message{
			Role:    schema.User,
			Content: prompt,
		}

		// 异步插入普通记忆
		go InsertMemory(ctx, conversationId, string(schema.User), prompt)

		return message, nil
	}
}

func GetHistoryMessageList(ctx context.Context, conversationId string, userId int64, prompt string) ([]adk.Message, string, error) {
	var err error
	var messages []adk.Message
	log.Printf("GetHistoryMessageList, conversationId: %s, userId: %d, prompt: %s", conversationId, userId, prompt)
	if conversationId == "" {
		conversationId, err = CreateConversation(ctx, userId)
		if err != nil {
			return nil, conversationId, err
		}
		go func() {
			resp, err := config.DefaultArkModel.Generate(ctx, []*schema.Message{
				{
					Role:    schema.User,
					Content: fmt.Sprintf("下面是用户的消息，请直接输出一个总结性的标题：%s", prompt),
				},
			})
			if err != nil {
				log.Printf("failed to generate title, err: %v", err)
				return
			}
			UpdateConversationTitle(ctx, conversationId, resp.Content)
		}()
	} else {
		chatMemory, err := GetMemoryList(ctx, conversationId)
		if err != nil {
			return nil, conversationId, err
		}

		for i, memory := range chatMemory {
			if memory.Metadata == "img-text" {
				// 查找对应的img-urls记录
				var imgUrls []string
				if chatMemory[i+1].Metadata == "img-urls" {
					// 解析JSON格式的img-urls
					if err := json.Unmarshal([]byte(chatMemory[i+1].Prompt), &imgUrls); err != nil {
						log.Printf("failed to unmarshal img-urls: %v", err)
						// 如果解析失败，作为普通文本消息处理
						messages = append(messages, &schema.Message{
							Role:    TransformMemoryRoleToMessage(memory.Type),
							Content: memory.Prompt,
						})
						continue
					}

					// 构造MultiContent消息
					multiContent := []schema.ChatMessagePart{
						{
							Type: schema.ChatMessagePartTypeText,
							Text: memory.Prompt,
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

					messages = append(messages, &schema.Message{
						Role:         TransformMemoryRoleToMessage(memory.Type),
						MultiContent: multiContent,
					})
				}
			} else if memory.Metadata == "img-urls" {
				// 跳过img-urls记录，因为它们已经在img-text处理中被使用
				continue
			} else {
				messages = append(messages, &schema.Message{
					Role:    TransformMemoryRoleToMessage(memory.Type),
					Content: memory.Prompt,
				})
			}
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
