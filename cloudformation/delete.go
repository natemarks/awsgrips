package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// DeleteStack Given a stack name, delete the stack
func DeleteStack(stackName string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	CFNClient := *cloudformation.NewFromConfig(cfg)

	input := &cloudformation.DeleteStackInput{StackName: aws.String(stackName)}
	_, err = CFNClient.DeleteStack(context.TODO(), input)
	return err
}
