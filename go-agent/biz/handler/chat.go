package handler

import (
	"context"
	"travel/biz/param"
	"travel/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetUploadUrlHandler(ctx context.Context, c *app.RequestContext) {
	var request param.UploadFileRequest
	if err := c.BindAndValidate(&request); err != nil {
		c.JSON(consts.StatusOK, param.Response{
			Code: consts.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	uploadUrl, err := service.NewChatService(ctx, c).GetUploadUrl(&request)
	if err != nil {
		c.JSON(consts.StatusOK, param.Response{
			Code: consts.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, param.Response{
		Code: consts.StatusOK,
		Msg:  "success",
		Data: uploadUrl,
	})
}