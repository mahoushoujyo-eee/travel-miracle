package config

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

var Retriever *milvus.Retriever

func InitRetriever(ctx context.Context) {
	// Create a retriever
	var err error
	Retriever, err = milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:    VDateBase,
		Embedding: ArkEmbeddingModel,
	})
	if err != nil {
		log.Fatalf("Failed to create retriver: %v", err)
		return
	}
	log.Printf("Retriver created success")
}