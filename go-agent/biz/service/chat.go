package service

import (
	"context"
	"errors"
	"log"
	"travel/biz/agent"
	"travel/biz/param"
	"travel/biz/util"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spf13/viper"
)


type ChatService struct {
	ctx context.Context
	c *app.RequestContext
}

func NewChatService(ctx context.Context, c *app.RequestContext) *ChatService {
	return &ChatService{
		ctx: ctx,
		c: c,
	}
}

func (s *ChatService) Chat(request *param.ChatRequest) (chan string, error) {
	conversationId := request.ConversationId
	var messages []adk.Message
	var err error
	if conversationId == ""{
		conversationId, err = agent.CreateConversation(s.ctx, request.UserId)
		if err != nil {
			return nil, err
		}
	}else{
		chatMemory, err := agent.GetMemoryList(s.ctx, conversationId)
		if err != nil {
			return nil, err
		}

		for _, memory := range chatMemory {
			messages = append(messages, &schema.Message{
				Role: schema.RoleType(memory.Type),
				Content: memory.Prompt,
			})
		}
	}

	if len(request.ImgUrls) > 0 {
		multiContent := []schema.ChatMessagePart{
			{
				Type: schema.ChatMessagePartTypeText,
				Text: request.Prompt,
			},
		}
		for _, url := range request.ImgUrls {
			multiContent = append(multiContent, schema.ChatMessagePart{
				Type: schema.ChatMessagePartTypeImageURL,
				ImageURL: &schema.ChatMessageImageURL{
					URL:    url,
					Detail: schema.ImageURLDetailAuto,
				},
			})
		}

		messages = append(messages, &schema.Message{
			Role: schema.User,
			MultiContent: multiContent,
		})
	}else{
		messages = append(messages, &schema.Message{
			Role: schema.User,
			Content: request.Prompt,
		})
	}

	iterator := agent.DefaultPlanRunner.Run(s.ctx, messages)
	
	for{
		event, ok := iterator.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			// 记录错误但不终止程序，允许继续处理
			log.Printf("\n事件处理错误: %v\n", event.Err)
		}

		if event.Output.MessageOutput.IsStreaming {
			// TODO: 处理流式输出
			continue
		}else{
			go agent.InsertMemory(s.ctx, conversationId, event.Output.MessageOutput)

		}
	}

	return nil, nil
}

func (s *ChatService) GetUploadUrl(request *param.UploadFileRequest) (*oss.PresignResult, error) {
	var ossRequest *param.GetUploadUrlRequest
	if request.Type == "image"{
		ossRequest = &param.GetUploadUrlRequest{
			Bucket: viper.GetString("oss.img-bucket"),
			Key: request.FileName,
			ContentType: request.ContentType,
		}
	}else if request.Type == "file"{
		ossRequest = &param.GetUploadUrlRequest{
			Bucket: viper.GetString("oss.file-bucket"),
			Key: request.FileName,
			ContentType: request.ContentType,
		}
	}else{
		return nil, errors.New("invalid upload type")
	}

	return util.GetUploadUrl(ossRequest, s.ctx)
}