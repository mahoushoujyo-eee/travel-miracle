package handler

import (
	"context"
	"travel/biz/param"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ServiceFunc 定义服务函数的类型，接收请求参数并返回响应数据和错误
type ServiceFunc[T any, R any] func(ctx context.Context, c *app.RequestContext, request *T) (*R, error)

// GenericHandler 通用的泛型处理器
// T: 请求参数类型
// R: 响应数据类型
func GenericHandler[T any, R any](serviceFunc ServiceFunc[T, R]) func(ctx context.Context, c *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		// 1. 绑定和验证请求参数
		var request T
		if err := c.BindAndValidate(&request); err != nil {
			c.JSON(consts.StatusOK, param.ResponseError(
				consts.StatusBadRequest,
				err.Error(),
			))
			return
		}

		// 2. 调用服务层处理业务逻辑
		result, err := serviceFunc(ctx, c, &request)
		if err != nil {
			c.JSON(consts.StatusOK, param.ResponseError(
				consts.StatusInternalServerError,
				err.Error(),
			))
			return
		}

		// 3. 返回成功响应
		c.JSON(consts.StatusOK, param.ResponseSuccess(result))
	}
}