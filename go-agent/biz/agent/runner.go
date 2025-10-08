package agent

import (
	"context"

	"github.com/cloudwego/eino/adk"
)



var(
	DefaultRunner *adk.Runner
)

func InitRunner(ctx context.Context){
    DefaultRunner = adk.NewRunner(ctx, adk.RunnerConfig{
       Agent:           DefaultChatModelAgent,
       CheckPointStore: NewInMemoryStore(),
       EnableStreaming: true,
    })
}

func NewRunner(ctx context.Context, agent adk.Agent) *adk.Runner {
    return adk.NewRunner(ctx, adk.RunnerConfig{
       Agent:           agent,
       CheckPointStore: NewInMemoryStore(),
       EnableStreaming: true,
    })
}