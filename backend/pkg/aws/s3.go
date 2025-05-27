package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

func NewS3Client(bucket string) (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("AWS設定の読み込みに失敗: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
		bucket: bucket,
	}, nil
}

func (s *S3Client) GetBlogPosts(ctx context.Context) ([]map[string]interface{}, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: &s.bucket,
		Prefix: aws.String("blogs/"),
	}

	result, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("S3オブジェクトの一覧取得に失敗: %v", err)
	}

	var posts []map[string]interface{}
	for _, item := range result.Contents {
		if *item.Key == "blogs/" {
			continue
		}

		getObjInput := &s3.GetObjectInput{
			Bucket: &s.bucket,
			Key:    item.Key,
		}

		result, err := s.client.GetObject(ctx, getObjInput)
		if err != nil {
			return nil, fmt.Errorf("オブジェクトの取得に失敗: %v", err)
		}

		var post map[string]interface{}
		if err := json.NewDecoder(result.Body).Decode(&post); err != nil {
			return nil, fmt.Errorf("JSONのデコードに失敗: %v", err)
		}

		posts = append(posts, post)
	}

	return posts, nil
} 