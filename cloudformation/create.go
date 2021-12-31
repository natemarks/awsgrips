package cloudformation

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/natemarks/awsgrips/fixtures"
)

func CreateStackInputFromLocalFile(filePath string) (result cloudformation.CreateStackInput, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	templateBody := string(content)
	if err != nil {
		return
	}
	result.TemplateBody = aws.String(templateBody)
	return
}

func CreateStackInputFromHTTP(httpUrl string) (result cloudformation.CreateStackInput, err error) {
	localFile, err := fixtures.DownloadUrl(httpUrl)
	if err != nil {
		return
	}
	result, err = CreateStackInputFromLocalFile(localFile)
	return
}

// CreateStackInputFromS3 Given a http url to template file in an S3 bucket,
// return a new cloudformation.CreateStackInput with the templateUrl set
// ex. https://s3.amazonaws.com/com.imprivata.709310380790.us-east-1.cloudformation-templates/topic.json
func CreateStackInputFromS3(s3Url string) (result cloudformation.CreateStackInput, err error) {
	result.TemplateURL = aws.String(s3Url)
	return
}

// CreateStack Create a stack and return the stackId
func CreateStack(input *cloudformation.CreateStackInput) (stackId string, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := *cloudformation.NewFromConfig(cfg)

	output, err := client.CreateStack(context.TODO(), input)
	if err != nil {
		return
	}
	stackId = *output.StackId
	return
}

// CreateStackWait Wait up to maxWaitInMinutes for the creation of a given stack (by stack name) to complete
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
