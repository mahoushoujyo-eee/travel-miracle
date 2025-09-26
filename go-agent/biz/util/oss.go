package util

import (
	"context"
	"log"
	"time"
	"travel/biz/config"
	"travel/biz/param"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

func GetUploadUrl(req *param.GetUploadUrlRequest, ctx context.Context) (*oss.PresignResult, error) {
	// 生成PutObject的预签名URL
	result, err := config.OssClient.Presign(ctx, &oss.PutObjectRequest{
		Bucket:      oss.Ptr(req.Bucket),
		Key:         oss.Ptr(req.Key),
		ContentType: oss.Ptr(req.ContentType), // 请确保在服务端生成该签名URL时设置的ContentType与在使用URL时设置的ContentType一致

	},
		oss.PresignExpires(10*time.Minute),
	)
	if err != nil {
		log.Printf("failed to put object presign %v", err)
		return nil, err
	}

	log.Printf("request method:%v\n", result.Method)
	log.Printf("request expiration:%v\n", result.Expiration)
	log.Printf("request url:%v\n", result.URL)

	return result, nil
}

func GetDownloadUrl(req *param.GetDownloadUrlRequest, ctx context.Context) (*oss.PresignResult, error) {
		// 生成PutObject的预签名URL
	result, err := config.OssClient.Presign(ctx, &oss.GetObjectRequest{
		Bucket:      oss.Ptr(req.Bucket),
		Key:         oss.Ptr(req.Key),
	},
		oss.PresignExpires(10*time.Minute),
	)
	if err != nil {
		log.Printf("failed to put object presign %v", err)
		return nil, err
	}

	log.Printf("request method:%v\n", result.Method)
	log.Printf("request expiration:%v\n", result.Expiration)
	log.Printf("request url:%v\n", result.URL)

	return result, nil
}