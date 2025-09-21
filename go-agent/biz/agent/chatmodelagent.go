package agent

import (
	"context"
	"fmt"
	"log"
	"travel/biz/config"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
)

var (
	DefaultChatModelAgent adk.Agent
)

func InitChatModelAgent(ctx context.Context){
	var err error
    DefaultChatModelAgent, err = adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
       Name:        "BookRecommender",
       Description: "An agent that can recommend books",
       Instruction: `You are an expert book recommender. Based on the user's request, use the "search_book" tool to find relevant books. Finally, present the results to the user.`,
       Model:       config.DefaultArkModel,
       ToolsConfig: adk.ToolsConfig{
          ToolsNodeConfig: compose.ToolsNodeConfig{
             Tools: config.ToolMap["fetch"],
          },
       },
    })
    if err != nil {
       log.Fatal(fmt.Errorf("failed to create chatmodel: %w", err))
    }
}

func NewAgent(ctx context.Context, name string, description string, instruction string, tools []tool.BaseTool) adk.Agent {
	agent,err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        name,
		Description: description,
		Instruction: instruction,
		Model:       config.DefaultArkModel,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: tools,
			},
		},
	})
	if err!= nil {
		log.Fatal(fmt.Errorf("failed to create chatmodel: %w", err))
	}
	return agent
}
