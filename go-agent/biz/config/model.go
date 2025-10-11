package config

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/model/openai"
	embeddingArk "github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/spf13/viper"
)

var(
	DefaultArkModel *ark.ChatModel
	DefaultMSModel 	*openai.ChatModel
	ArkEmbeddingModel *embeddingArk.Embedder
	DefaultVisionModel *ark.ChatModel
)

func InitModel(ctx context.Context) {
	log.Printf("Initalizing model...")

	timeout := 120 * time.Second  // 增加超时时间到120秒

	var err error
	DefaultArkModel, err = ark.NewChatModel(ctx, &ark.ChatModelConfig{
		// 服务配置
		BaseURL: viper.GetString("llm-ark.base-url"), // 服务地址
		Region:  "cn-beijing",                               // 区域
		APIKey: viper.GetString("llm-ark.api-key"), // API Key 认证
		// 模型配置
		Model:   viper.GetString("llm-ark.default-model"), // 模型端点 ID
		Timeout: &timeout,
	})

	if err != nil {
		log.Fatalf("failed to init chat model, err: %v", err)
	}

	DefaultMSModel, err = openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: viper.GetString("llm-ms.base-url"),
		APIKey:  viper.GetString("llm-ms.api-key"),
		Model:   viper.GetString("llm-ms.default-model"),
		Timeout: timeout,
	})
	
	if err != nil {
		log.Fatalf("failed to init chat model, err: %v", err)
	}

	ArkEmbeddingModel, err = embeddingArk.NewEmbedder(ctx, &embeddingArk.EmbeddingConfig{
        APIKey:  viper.GetString("embedding-ark.api-key"),
        Model:   viper.GetString("embedding-ark.default-model"),
        Timeout: &timeout,
    })
	if err != nil {
		log.Fatalf("failed to init embedding model, err: %v", err)
	}

	DefaultVisionModel, err = ark.NewChatModel(ctx, &ark.ChatModelConfig{
		// 服务配置
		BaseURL: viper.GetString("llm-vision-ark.base-url"), // 服务地址
		Region:  "cn-beijing",                               // 区域
		APIKey: viper.GetString("llm-vision-ark.api-key"), // API Key 认证
		// 模型配置
		Model:   viper.GetString("llm-vision-ark.default-model"), // 模型端点 ID
		Timeout: &timeout,
	})
	if err != nil {
		log.Fatalf("failed to init vision model, err: %v", err)
	}
	
	log.Printf("Model initalized!")
}