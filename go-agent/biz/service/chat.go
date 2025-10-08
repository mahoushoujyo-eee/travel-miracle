package service

import (
	"context"
	"errors"
	"log"
	"travel/biz/agent"
	"travel/biz/param"
	"travel/biz/util"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spf13/viper"
)

type ChatService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewChatService(ctx context.Context, c *app.RequestContext) *ChatService {
	return &ChatService{
		ctx: ctx,
		c:   c,
	}
}

func (s *ChatService) Chat(request *param.ChatRequest, responseChan chan *param.SSEChatResponse) error {
	// 添加panic恢复机制，防止程序崩溃
	defer s.handleChatPanic(request, responseChan)

	conversationId := request.ConversationId
	messages, conversationId, err := agent.GetHistoryMessageList(s.ctx, conversationId, request.UserId, request.Prompt)
	if err != nil {
		return err
	}
	responseChan <- &param.SSEChatResponse{
		Type:           "start",
		ConversationId: conversationId,
	}

	// 使用封装的函数创建用户消息
	userMessage, err := agent.CreateUserMessage(s.ctx, conversationId, request.Prompt, request.ImgUrls)
	if err != nil {
		return err
	}
	messages = append(messages, userMessage)

	iterator := agent.DefaultPlanRunner.Run(s.ctx, messages)

	for {
		event, ok := iterator.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			// 记录错误但不终止程序，允许继续处理
			log.Printf("\n事件处理错误: %v\n", event.Err)
			continue // 跳过这个事件，继续处理下一个
		}

		// 添加nil检查，防止nil pointer dereference
		if event.Output == nil || event.Output.MessageOutput == nil {
			log.Printf("\n事件输出为空，跳过处理\n")
			continue
		}

		if event.Output.MessageOutput.IsStreaming {
			// TODO: 处理流式输出
			content := ""
			reasoningContent := ""
			stream := event.Output.MessageOutput.MessageStream

			// 检查stream是否为nil
			if stream == nil {
				log.Printf("\n流式传输stream为空，跳过处理\n")
				continue
			}

			for {
				msg, err := stream.Recv()
				if err != nil {
					// 检查是否是正常结束或可恢复的错误
					if err.Error() == "EOF" || msg == nil {
						log.Printf("\n流式传输正常结束\n")
						break
					}
					// 对于超时等网络错误，记录日志但不终止程序
					log.Printf("\n流式传输错误: %v\n", err)
					break
				}
				if msg == nil {
					break
				}

				if msg.ReasoningContent != "" {
					reasoningContent += msg.ReasoningContent
					response := &param.SSEChatResponse{
						Type:           "stream-reasoning",
						Content:        msg.ReasoningContent,
						ConversationId: conversationId,
					}
					responseChan <- response
				}
				if msg.Content != "" {
					content += msg.Content
					response := &param.SSEChatResponse{
						Type:           "stream-chat",
						Content:        msg.Content,
						ConversationId: conversationId,
					}
					responseChan <- response
				}
			}

			go func() {
				if reasoningContent != "" {
					agent.InsertMemory(s.ctx, conversationId, "stream-reasoning", reasoningContent)
				}
				if content != "" {
					agent.InsertMemory(s.ctx, conversationId, "stream-chat", content)
				}
			}()

			continue
		} else {
			var response *param.SSEChatResponse

			// 添加额外的nil检查，防止访问Message时出现panic
			if event.Output.MessageOutput.Message == nil {
				log.Printf("\n消息内容为空，跳过处理\n")
				continue
			}

			if event.Output.MessageOutput.Message.ToolName != "" {
				go agent.InsertMemoryWithTool(s.ctx, conversationId, string(event.Output.MessageOutput.Role), event.Output.MessageOutput.Message.Content, event.Output.MessageOutput.Message.ToolName)
				response = &param.SSEChatResponse{
					Type:           string(event.Output.MessageOutput.Role) + ":" + event.Output.MessageOutput.Message.ToolName,
					Content:        event.Output.MessageOutput.Message.Content,
					ConversationId: conversationId,
				}
				log.Printf("\n工具调用: %v\n", event.Output.MessageOutput.Message.ToolCalls)
			} else {
				go agent.InsertMemory(s.ctx, conversationId, string(event.Output.MessageOutput.Role), event.Output.MessageOutput.Message.Content)
				response = &param.SSEChatResponse{
					Type:           string(event.Output.MessageOutput.Role),
					Content:        event.Output.MessageOutput.Message.Content,
					ConversationId: conversationId,
				}
			}
			responseChan <- response
		}
	}

	return nil
}

func (s *ChatService) GetUploadUrl(request *param.UploadFileRequest) (*oss.PresignResult, error) {
	var ossRequest *param.GetUploadUrlRequest
	switch request.Type {
	case "image":
		ossRequest = &param.GetUploadUrlRequest{
			Bucket:      viper.GetString("oss.img-bucket"),
			Key:         request.FileName,
			ContentType: request.ContentType,
		}
	case "file":
		ossRequest = &param.GetUploadUrlRequest{
			Bucket:      viper.GetString("oss.file-bucket"),
			Key:         request.FileName,
			ContentType: request.ContentType,
		}
	default:
		return nil, errors.New("invalid upload type")
	}

	return util.GetUploadUrl(ossRequest, s.ctx)
}

// handleChatPanic 处理Chat函数中的panic恢复
func (s *ChatService) handleChatPanic(request *param.ChatRequest, responseChan chan *param.SSEChatResponse) {
	if r := recover(); r != nil {
		log.Printf("Chat函数发生panic，已恢复: %v", r)
		// 发送错误响应给客户端
		errorResponse := &param.SSEChatResponse{
			Type:           "error",
			Content:        "服务暂时不可用，请稍后重试",
			ConversationId: request.ConversationId,
		}
		select {
		case responseChan <- errorResponse:
		default:
			// 如果channel已关闭，忽略
		}
	}
}
