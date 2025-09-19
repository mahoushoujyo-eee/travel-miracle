package config

import (
	"context"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/spf13/viper"
)

var VDateBase client.Client

func InitVDateBase(ctx context.Context) {
	// Get the environment variables
	addr := viper.GetString("milvus.addr")
	
	// Create a client
	var err error
	VDateBase, err = client.NewClient(ctx, client.Config{
		Address:  addr,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return
	}
}