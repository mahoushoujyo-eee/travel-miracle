package service

import (
	"context"
	"errors"
	"travel/biz/param"
	"travel/biz/util"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
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