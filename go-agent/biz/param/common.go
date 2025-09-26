package param

import "github.com/cloudwego/hertz/pkg/protocol/consts"

type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(data interface{}) Response {
	return Response{
		Code: consts.StatusOK,
		Msg:  "success",
		Data: data,
	}
}

func ResponseError(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}