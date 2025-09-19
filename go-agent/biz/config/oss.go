package config

import (
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/spf13/viper"
)

var (
	OssClient *oss.Client
)

func InitOssClient(ctx context.Context){
	config := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(viper.GetString("oss.region"))

	OssClient = oss.NewClient(config)
}