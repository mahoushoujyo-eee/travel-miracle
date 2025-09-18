package config

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/mark3labs/mcp-go/client"
	goMcp "github.com/mark3labs/mcp-go/mcp"
)

var Tools = []tool.BaseTool{}
var ToolNode *compose.ToolsNode

type MCPServer struct {
	URL string
	Type string
}

var servers = []MCPServer{
	{URL: "https://mcp.api-inference.modelscope.net/9771b53107984b/mcp", Type: "shttp"},
	{URL: "https://mcp.api-inference.modelscope.net/b1ea70ecdcba49/mcp", Type: "shttp"},
	{URL: "https://mcp.api-inference.modelscope.net/a6b39c63b2944e/mcp", Type: "shttp"},

}

func InitMcpTools(ctx context.Context) {
	for _, server := range servers {
		var cli client.MCPClient
		var err error
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

		Tools = append(Tools, mcpTools...)
		log.Printf("successfully initialized MCP server: %s, tools count: %d", server.URL, len(mcpTools))
	}

	log.Printf("mcp tools initialized, total count: %d", len(Tools))
	
	if len(Tools) == 0 {
		log.Printf("no MCP tools available, skipping tool node creation")
		return
	}
	
	var err error
	ToolNode, err = compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: Tools,
	})
	if err != nil {
		log.Printf("failed to create tool node: %v", err)
		return
	}

	for _, t := range Tools {
		info, err := t.Info(ctx)
		if err != nil {
			panic(err)
		}
		log.Printf("tool name: %s, desc: %s", info.Name, info.Desc)
	}
}


