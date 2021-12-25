package cloudformation

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"io/ioutil"
	"time"
)

// CreateStack restore a given snapshot to a given instnace namd
func CreateStack(stackName string, stackFile string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	CFNClient := *cloudformation.NewFromConfig(cfg)
	content, err := ioutil.ReadFile(stackFile)
	templateBody := string(content)
	if err != nil {
		return "", err
	}
	input := &cloudformation.CreateStackInput{StackName: aws.String(stackName),
		TemplateBody: aws.String(templateBody)}
	output, err := CFNClient.CreateStack(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return *output.StackId, err
}

func CreateStackWait(stackName string, maxWaitInMinutes int) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := cloudformation.NewFromConfig(cfg)
	waiter := cloudformation.NewStackCreateCompleteWaiter(client)
	params := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}
	maxWaitTime := time.Duration(maxWaitInMinutes) * time.Minute
	err = waiter.Wait(context.TODO(), params, maxWaitTime)
	return err
}
