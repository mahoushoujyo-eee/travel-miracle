package handler

import (
	"context"
	"travel/biz/param"
	"travel/biz/service"
	"travel/biz/util"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sse"
)

// GetUploadUrlHandler 使用通用泛型处理器
var GetUploadUrlHandler = GenericHandler(
	func(ctx context.Context, c *app.RequestContext, request *param.UploadFileRequest) (*oss.PresignResult, error) {
		return service.NewChatService(ctx, c).GetUploadUrl(request)
	},
)

func ChatHandler(ctx context.Context, c *app.RequestContext) {
	request := new(param.ChatRequest)
	if err := c.BindAndValidate(request); err != nil {
		c.JSON(consts.StatusOK, param.ResponseError(consts.StatusInternalServerError, err.Error()))
		return
	}
	responseChan := make(chan *param.SSEChatResponse)
	sseSender := util.NewSSESender(sse.NewStream(c))

	go func() {
		defer close(responseChan)
		if err := service.NewChatService(ctx, c).Chat(request, responseChan); err != nil {
			c.JSON(consts.StatusOK, param.ResponseError(consts.StatusInternalServerError, err.Error()))
			return
		}
	}()
	for response := range responseChan {
		sseSender.Send(ctx, &sse.Event{
			Event: response.Type,
			Data:  []byte(response.Content),
		})
	}
}
