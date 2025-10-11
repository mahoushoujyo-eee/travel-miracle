package agent

import (
	"context"
	"fmt"
	"log"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
)

var (
	DefaultPlanAgent adk.Agent
	DefaultPlanRunner *adk.Runner
	DefaultRecommendAgent adk.Agent
	DefaultRecommendRunner *adk.Runner
)


func InitPlanRunner(ctx context.Context){
	// 合并多个工具
	var allTools []tool.BaseTool
	allTools = append(allTools, config.ToolMap["amap"]...)    // 高德地图工具
	allTools = append(allTools, config.ToolMap["12306"]...)  // 12306工具

	name := "出行路线规划小助手"
	description := "一个帮助他人规划出行路线的助手"
	systemPrompt := "你是一个可以使用高德地图和12306工具帮助他人规划出行路线的助手，可以查询地图信息和火车票信息，根据他人的喜好来推荐多条出行路线"
	DefaultPlanAgent = NewAgent(ctx, name, description, systemPrompt, allTools)
    DefaultPlanRunner = NewRunner(ctx, DefaultPlanAgent)
}

func InitRecommendRunner(ctx context.Context){
	// 合并多个工具
	var allTools []tool.BaseTool
	allTools = append(allTools, config.ToolMap["jina"]...)    // jina工具
	allTools = append(allTools, config.ToolMap["fetch"]...)  // fetch工具

	name := "景点推荐小助手"
	description := "一个帮助他人推荐景点的助手"
	systemPrompt := "你是一个可以使用jina和fetch工具帮助他人推荐景点的助手，可以搜索和获取景点信息，根据他人的喜好来推荐合适的景点。"
	DefaultRecommendAgent = NewAgent(ctx, name, description, systemPrompt, allTools)
    DefaultRecommendRunner = NewRunner(ctx, DefaultRecommendAgent)
}

func HandleIterator(iterator *adk.AsyncIterator[*adk.AgentEvent]) {
	for {
		event, ok := iterator.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			// 记录错误但不终止程序，允许继续处理
			fmt.Printf("\n事件处理错误: %v\n", event.Err)
		}

		if event.Output.MessageOutput.IsStreaming {
			stream := event.Output.MessageOutput.MessageStream
			for {
				msg, err := stream.Recv()
				if err != nil {
					// 检查是否是正常结束或可恢复的错误
					if err.Error() == "EOF" || msg == nil {
						fmt.Printf("\n流式传输正常结束\n")
						break
					}
					// 对于超时等网络错误，记录日志但不终止程序
					fmt.Printf("\n流式传输错误: %v\n", err)
					break
				}
				if msg == nil {
					break
				}
				if msg.Content != ""{
					fmt.Printf("%v", msg)
				}

				if msg.ReasoningContent != "" {
					fmt.Printf("%s", msg.ReasoningContent)
				}
			}
		}else{
			msg, err := event.Output.MessageOutput.GetMessage()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nmessage:\n%v\n", msg.Content)
		}
	}
}