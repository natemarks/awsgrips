package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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

// ListObjects given a bucket and a prefix string return a (recursive) list of contents
// ListObjectsV2 returns in chunks of 1000
// loop on the continuationToken to get everything
// until the resp.IsTruncated comes back as False
// https://stackoverflow.com/questions/65043012/how-to-get-more-than-1000-objects-from-s3
func ListObjects(bucket, prefix string) (results []types.Object, err error) {
	var continuationToken *string
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return results, err
	}

	client := *s3.NewFromConfig(cfg)

	// loop using continuationToken becuase ListObjectsV2 is limited to 1000
	for {
		resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket:            aws.String(bucket),
			Prefix:            aws.String(prefix),
			ContinuationToken: continuationToken,
		})
		if err != nil {
			return results, err
		}
		results = append(results, resp.Contents...)
		if !resp.IsTruncated {
			break
		}
		continuationToken = resp.NextContinuationToken
	}
	return results, err
}
