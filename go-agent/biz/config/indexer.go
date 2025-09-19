package config

import (
	"context"
	"log"

	"github.com/cloudwego/eino/schema"

	"github.com/cloudwego/eino-ext/components/indexer/milvus"
)

var Indexer *milvus.Indexer		

func InitIndexer(ctx context.Context) {
	// Create an indexer
	var err error
	Indexer, err = milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:    VDateBase,
		Embedding: ArkEmbeddingModel,
	})
	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
		return
	}
	log.Printf("Indexer created success")
}

func RunDemo(ctx context.Context){
		// Store documents
	docs := []*schema.Document{
		{
			ID:      "milvus-1",
			Content: "milvus is an open-source vector database",
			MetaData: map[string]any{
				"h1": "milvus",
				"h2": "open-source",
				"h3": "vector database",
			},
		},
		{
			ID:      "milvus-2",
			Content: "milvus is a distributed vector database",
		},
	}
	ids, err := Indexer.Store(ctx, docs)
	if err != nil {
		log.Fatalf("Failed to store: %v", err)
		return
	}
	log.Printf("Store success, ids: %v", ids)
}
