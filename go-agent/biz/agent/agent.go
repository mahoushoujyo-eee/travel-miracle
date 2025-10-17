package agent

import (
	"context"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
)

var (
	DefaultPlanAgent adk.Agent
	DefaultPlanRunner *adk.Runner
	DefaultRecommendAgent adk.Agent
	DefaultRecommendRunner *adk.Runner
	DefaultWebSearchAgent adk.Agent
	DefaultWebSearchRunner *adk.Runner
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
	// allTools = append(allTools, config.ToolMap["jina"]...)    // jina工具
	// allTools = append(allTools, config.ToolMap["fetch"]...)  // fetch工具
	allTools = append(allTools, config.ToolMap["amap"]...)    // 高德地图工具

	name := "景点推荐小助手"
	description := "一个帮助他人推荐景点的助手"
	systemPrompt := "你是一个可以使用工具帮助他人推荐景点的助手，在推荐周围景点的时候先使用amap相关工具来搜索周围的景点，然后再调用搜索工具查看景点详细信息，最后根据用户的喜好来推荐合适的景点，注意如果无特殊要求请不要推荐大于十个景点。"
	DefaultRecommendAgent = NewAgent(ctx, name, description, systemPrompt, allTools)
    DefaultRecommendRunner = NewRunner(ctx, DefaultRecommendAgent)
}

func InitWebSearchRunner(ctx context.Context){
	// 合并多个工具
	var allTools []tool.BaseTool
	allTools = append(allTools, config.ToolMap["jina"]...)    // jina工具
	allTools = append(allTools, config.ToolMap["fetch"]...)  // fetch工具

	name := "网络搜索小助手"
	description := "一个帮助他人进行网络搜索的助手"
	systemPrompt := "你是一个可以使用工具帮助他人进行网络搜索的助手，在搜索的时候先使用jina\fetch相关工具来搜索用户需要的内容并进行一定总结。"
	DefaultWebSearchAgent = NewAgent(ctx, name, description, systemPrompt, allTools)
    DefaultWebSearchRunner = NewRunner(ctx, DefaultWebSearchAgent)
}