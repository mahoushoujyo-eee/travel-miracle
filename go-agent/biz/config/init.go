package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitAll(ctx context.Context) {
	loadEnv()
	loadViper()
	InitDatabase(ctx)
	InitOssClient(ctx)
	InitModel(ctx)
	InitMcpTools(ctx)
	InitRedisClient(ctx)
	InitCozeloop(ctx)
	// config.InitVDateBase(ctx)
	// config.InitRetriever(ctx)
	// config.InitIndexer(ctx)
}

func loadViper() {
	// configure file name
	viper.SetConfigName("config")
	// configure file type
	viper.SetConfigType("yaml")
	// configure viper to read from the current directory
	viper.AddConfigPath(".") // current directory
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func loadEnv() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}