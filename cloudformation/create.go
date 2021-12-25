package cloudformation

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"io/ioutil"
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
	return *output.StackId, err
}
