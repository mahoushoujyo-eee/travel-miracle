package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"travel/biz/agent"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func main(){
	ctx := context.Background()
	var allTools []tool.BaseTool
	config.InitAll(ctx)

	allTools = append(allTools, config.ToolMap["jina"]...)
	allTools = append(allTools, config.ToolMap["fetch"]...)

	myAgent := agent.NewAgent(ctx, 
		"景点推荐小助手", 
		"根据所给地址推荐景点的助手",
		fmt.Sprintf("你是一个给用户推荐景点的助手，善于使用工具来联网查找用户所给地点的可游玩景点，并且要考虑当前月份是%v，最后生成总结输出。注意：要在搜索两次详情页后总结内容，节省调用的开销。", time.Now()), 
		allTools)

	myRunner := agent.NewRunner(ctx, myAgent)
	iter := myRunner.Run(ctx, []adk.Message{
		{
			Role:    schema.User,
			Content: "我大概下周想要去南京旅游，给我推荐下可以游玩的景点以及他们的详细资料呗",
		},
	})
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			log.Fatal(event.Err)
		}
		if event.Output.MessageOutput.IsStreaming {
			stream := event.Output.MessageOutput.MessageStream
			for {
				msg, err := stream.Recv()
				if err != nil {
					fmt.Printf("error:\n%v\n======\n", err)
					continue
				}
				if msg == nil {
					continue
				}
				fmt.Printf("\nmessage:\n%v\n======", msg)
			}
		}else{
			msg, err := event.Output.MessageOutput.GetMessage()
			if err != nil {
				fmt.Printf("error:\n%v\n======\n", err)
				continue
			}
			fmt.Printf("message:\n%v\n======\n", msg)
		}
	}
}