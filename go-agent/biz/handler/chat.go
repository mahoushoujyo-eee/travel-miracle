package handler

import (
	"context"
	"travel/biz/param"
	"travel/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

// GetUploadUrlHandler 使用通用泛型处理器
var GetUploadUrlHandler = GenericHandler(
	func(ctx context.Context, c *app.RequestContext, request *param.UploadFileRequest) (*oss.PresignResult, error) {
		return service.NewChatService(ctx, c).GetUploadUrl(request)
	},
)
