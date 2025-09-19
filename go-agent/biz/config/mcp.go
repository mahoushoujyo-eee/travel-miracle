package config

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/compose"
	"github.com/mark3labs/mcp-go/client"
	goMcp "github.com/mark3labs/mcp-go/mcp"
)

var ToolNode *compose.ToolsNode
var ToolNodeMap = make(map[string]*compose.ToolsNode)

type MCPServer struct {
	URL string
	Type string
	Name string
}

var servers = []MCPServer{
	{URL: "https://mcp.api-inference.modelscope.net/9771b53107984b/mcp", Type: "shttp", Name:"fetch"},
	{URL: "https://mcp.api-inference.modelscope.net/4d2d99ac6a974a/mcp", Type: "shttp", Name:"bing"},
	{URL: "https://mcp.api-inference.modelscope.net/b1ea70ecdcba49/mcp", Type: "shttp", Name:"amp"},
	{URL: "https://mcp.api-inference.modelscope.net/7355509d0d7341/mcp", Type: "shttp", Name:"baidu-map"},
	{URL: "https://mcp.api-inference.modelscope.net/fd8a083b572f4c/mcp", Type: "shttp", Name:"12306"},
}

func InitMcpTools(ctx context.Context) {
	for _, server := range servers {
		var cli client.MCPClient
		var err error
		var toolNodes *compose.ToolsNode
		switch server.Type {
		case "sse":
			cli, err = client.NewSSEMCPClient(server.URL)
			//cli.Start(ctx) looks like need to execute this start method, but may have dependency version issue, so comment it
		case "shttp":
			cli, err = client.NewStreamableHttpClient(server.URL)
		default:
			log.Printf("unsupported server type: %s, url: %s", server.Type, server.URL)
			continue
		}
		if err != nil {
			log.Printf("create mcp client failed, url: %s, err: %v", server.URL, err)
			continue
		}

		initRequest := goMcp.InitializeRequest{}
		initRequest.Params.ProtocolVersion = goMcp.LATEST_PROTOCOL_VERSION
		initRequest.Params.ClientInfo = goMcp.Implementation{
			Name:    "client",
			Version: "1.0.0",
		}
		_, err = cli.Initialize(ctx, initRequest)
		if err != nil {
			log.Printf("initialize mcp client failed, url: %s, err: %v", server.URL, err)
			continue
		}

		mcpTools, err := mcp.GetTools(ctx, &mcp.Config{Cli: cli})
		if err != nil {
			log.Printf("get mcp tools failed, url: %s, err: %v", server.URL, err)
			continue
		}
		log.Printf("successfully initialized MCP server: %s, tools count: %d", server.URL, len(mcpTools))

		for _, t := range mcpTools {
			info, err := t.Info(ctx)
			if err != nil {
				log.Printf("failed to get tool info: %v", err)
				continue
			}
			log.Printf("tool name: %s, desc: %s", info.Name, info.Desc)
		}

		toolNodes, err = compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
			Tools: mcpTools,
		})
		if err != nil {
			log.Printf("failed to create tool node: %v", err)
			continue
		}
		ToolNodeMap[server.Name] = toolNodes
	}
}


