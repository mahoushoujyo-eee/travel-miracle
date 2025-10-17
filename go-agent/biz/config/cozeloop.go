package config

import (
	"context"
	"log"

	ccb "github.com/cloudwego/eino-ext/callbacks/cozeloop"
	"github.com/cloudwego/eino/callbacks"
	"github.com/coze-dev/cozeloop-go"
)


func InitCozeloop(ctx context.Context) {
	log.Println("Init cozeloop")
	client, err := cozeloop.NewClient()
	if err != nil {
		panic(err)
	}
	// 在服务 init 时 once 调用
	handler := ccb.NewLoopHandler(client)
	callbacks.AppendGlobalHandlers(handler)
	log.Println("Cozeloop init success")
}