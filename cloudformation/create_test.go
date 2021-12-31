package cloudformation

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

func TestCreateStack(t *testing.T) {
	type args struct {
		input *cloudformation.CreateStackInput
	}
	tests := []struct {
		name        string
		args        args
		wantStackId string
		wantErr     bool
	}{
		{name: "valid", args: args{input: &cloudformation.CreateStackInput{
			StackName:                   aws.String("TestCreateStack-valid"),
			Capabilities:                nil,
			ClientRequestToken:          nil,
			DisableRollback:             nil,
			EnableTerminationProtection: nil,
			NotificationARNs:            nil,
			OnFailure:                   "",
			Parameters:                  []types.Parameter{{ParameterKey: aws.String("BucketName"), ParameterValue: aws.String("natemarks-test-create-stack-valid-2874gf8b24byv")}},
			ResourceTypes:               nil,
			RoleARN:                     nil,
			RollbackConfiguration:       nil,
			StackPolicyBody:             nil,
			StackPolicyURL:              nil,
			Tags:                        []types.Tag{{Key: aws.String("deleteme"), Value: aws.String("true")}},
			TemplateBody:                nil,
			TemplateURL:                 aws.String("https://natemarks-cloudformation-public.s3.amazonaws.com/cfn-s3buckets/private.json"),
			TimeoutInMinutes:            nil,
		}},
			wantErr:     false,
			wantStackId: "TestCreateStack-valid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStackId, err := CreateStack(tt.args.input)
			_ = CreateStackWait(tt.wantStackId, 5)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !strings.Contains(gotStackId, tt.wantStackId) {
				t.Errorf("CreateStack() gotStackId does not contain stack name: %v", tt.wantStackId)
			}
		})
		_ = DeleteStack(tt.wantStackId)
	}
}

func TestCreateStackInputFromHTTP(t *testing.T) {
	type args struct {
		httpUrl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{httpUrl: "https://raw.githubusercontent.com/natemarks/cfn-s3buckets/main/private.json"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := CreateStackInputFromHTTP(tt.args.httpUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStackInputFromHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(*gotResult.TemplateBody, "AWSTemplateFormatVersion") {
				t.Error("TemplateBody doesn't contain expected string 'AWSTemplateFormatVersion'")
			}
		})
	}
}

func TestCreateStackInputFromS3(t *testing.T) {
	type args struct {
		s3Url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				s3Url: "https://natemarks-cloudformation-public.s3.amazonaws.com/cfn-s3buckets/private.json",
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := CreateStackInputFromS3(tt.args.s3Url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStackInputFromS3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(*gotResult.TemplateURL == tt.args.s3Url) {
				t.Error("TemplateURL is not the s3Url")
			}
		})
	}
}
