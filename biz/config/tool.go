package config

import (
	"context"
	"log"
	"time"
	
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/ddgsearch"
)

var (	
	DdgTool tool.InvokableTool
)

func InitTools(ctx context.Context) {
	var err error
	DdgTool, err = duckduckgo.NewTool(ctx, &duckduckgo.Config{
		MaxResults: 5, // Limit to return 3 results
		Region:     ddgsearch.RegionCN,
		DDGConfig: &ddgsearch.Config{
			Timeout:    10 * time.Second,
			Cache:      true,
			MaxRetries: 5,
		},
	})
	if err != nil {
		log.Fatalf("failed to init tools, err: %v", err)
	}
}
