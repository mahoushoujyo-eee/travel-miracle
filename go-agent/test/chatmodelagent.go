package main

import (
	"context"
	"fmt"
	"log"
	"travel/biz/agent"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	answer := "根据下面内容为我生成更规范的回复："
	config.InitAll(ctx)
	// 合并多个工具
	var allTools []tool.BaseTool
	allTools = append(allTools, config.ToolMap["amap"]...)    // 高德地图工具
	allTools = append(allTools, config.ToolMap["12306"]...)  // 12306工具
	allTools = append(allTools, config.ToolMap["baidu-map"]...)    // 百度地图工具

	
	myAgent := agent.NewAgent(ctx, "出行路线规划小助手", "一个帮助他人规划出行路线的助手", "你是一个可以使用地图工具和12306查火车高铁工具帮助他人规划出行路线的助手，可以查询地图信息和火车票信息，根据他人的喜好来推荐多条出行路线,如果没有提供偏好你可以考虑提供自驾、高铁、飞机等出行方式进行规划推荐", allTools)
	myRunner := agent.NewRunner(ctx, myAgent)
	iter := myRunner.Run(ctx, []adk.Message{
		{
			Role:    schema.User,
			Content: "我大概下周想要去威海旅游，为我规划下从南京到威海的出行路线，我目前考虑驾车和坐高铁两种方案",
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
					log.Fatal(err)
				}
				if msg == nil {
					break
				}
				fmt.Printf("\nmessage:\n%v\n======", msg)
				answer = answer + msg.Content
			}
		}else{
			msg, err := event.Output.MessageOutput.GetMessage()
			if err != nil {
				fmt.Printf("\nerror:\n%v\n======", err)
				continue
			}
			fmt.Printf("\nmessage:\n%v\n======", msg)
			answer = answer + msg.Content
		}
	}
	fmt.Printf("sum: %s", answer)
	fmt.Printf("\n\n")


	sumAgent := agent.NewAgent(ctx, "总结解释小助手", "根据已有的出行规划信息进行解释和总结的助手", "你是一个擅长总结和解释的助手，能够根据已有的出行规划信息，把内容形式变得更像一个人类回复，并且可以展开解释其中内容和概括其中内容", nil)
	sumRunner := agent.NewRunner(ctx, sumAgent)
	iter = sumRunner.Query(ctx, answer)
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
					log.Fatal(err)
				}
				if msg == nil {
					break
				}
				fmt.Printf("\nmessage:\n%v\n======", msg)
			}
		}else{
			msg, err := event.Output.MessageOutput.GetMessage()
			if err != nil {
				fmt.Printf("\nerror:\n%v\n======", err)
				continue
			}
			fmt.Printf("\nmessage:\n%v\n======", msg)
		}
	}

}


func main1() {
	ctx := context.Background()
	config.InitAll(ctx)
	// 合并多个工具
	var allTools []tool.BaseTool
	allTools = append(allTools, config.ToolMap["amp"]...)    // 高德地图工具
	allTools = append(allTools, config.ToolMap["12306"]...)  // 12306工具
	
	myAgent := agent.NewAgent(ctx, "出行路线规划小助手", "一个帮助他人规划出行路线的助手", "你是一个可以使用高德地图和12306工具帮助他人规划出行路线的助手，可以查询地图信息和火车票信息，根据他人的喜好来推荐多条出行路线", allTools)
	myRunner := agent.NewRunner(ctx, myAgent)
	iter := myRunner.Query(ctx, "请为我规划下从南京到潍坊的出行路线，谢谢！")
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			log.Fatal(event.Err)
		}
		msg, err := event.Output.MessageOutput.GetMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nmessage:\n%v\n======", msg)
	}
}
