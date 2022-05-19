// +build !unit
package fixtures

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func TestDownloadUrl(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid", args: args{location: "https://gist.githubusercontent.com/natemarks/cb6dab6de8adec96aea3cc9f4c0bca94/raw/aab5263b369f9a201cc6a99915d45a24e9f6e4a1/valid.json"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocalFile, err := DownloadUrl(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, err := os.Stat(gotLocalFile); err != nil {
				t.Errorf("DownloadUrl() Unable to stat downloaded file: %v ", gotLocalFile)
			}
		})
	}
}

func TestTempTagCreateStackInput(t *testing.T) {
	type args struct {
		input *cloudformation.CreateStackInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid",
			args: args{input: &cloudformation.CreateStackInput{
				StackName:                   aws.String("TestTempTagCreateStackInputValid"),
				Capabilities:                nil,
				ClientRequestToken:          nil,
				DisableRollback:             nil,
				EnableTerminationProtection: nil,
				NotificationARNs:            nil,
				OnFailure:                   "",
				Parameters:                  nil,
				ResourceTypes:               nil,
				RoleARN:                     nil,
				RollbackConfiguration:       nil,
				StackPolicyBody:             nil,
				StackPolicyURL:              nil,
				Tags:                        nil,
				TemplateBody:                nil,
				TemplateURL:                 nil,
				TimeoutInMinutes:            nil,
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TempTagCreateStackInput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("TempTagCreateStackInput() error = %v, wantErr %v", err, tt.wantErr)
			}
			numberOfTags := len(tt.args.input.Tags)
			if numberOfTags != 2 {
				t.Errorf("Expected 2 tags found: %d", numberOfTags)
			}
		})
	}
}
