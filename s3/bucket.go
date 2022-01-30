package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CreatePath given a bucket and a path string create the path
// make sure the last character in the path is a '/' to create a folder
func CreatePath(bucket, path string) (err error) {
	// make sure the last character is a "/"
	lastChar := path[len(path)-1:]
	if lastChar != "/" {
		path = path + "/"
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := *s3.NewFromConfig(cfg)
	input := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}
	_, err = client.PutObject(context.TODO(), &input)
	return err
}
