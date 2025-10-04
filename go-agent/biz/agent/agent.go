package agent

import (
	"context"
	"fmt"
	"time"
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
	allTools = append(allTools, config.ToolMap["amp"]...)    // 高德地图工具
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
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	systemPrompt := fmt.Sprintf("你是一个可以使用jina和fetch工具帮助他人推荐景点的助手，可以搜索和获取景点信息，根据他人的喜好来推荐合适的景点。当前时间：%s", currentTime)
	DefaultRecommendAgent = NewAgent(ctx, name, description, systemPrompt, allTools)
    DefaultRecommendRunner = NewRunner(ctx, DefaultRecommendAgent)
}