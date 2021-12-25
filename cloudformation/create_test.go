package cloudformation

import (
	"strings"
	"testing"
)

func TestCreateStack(t *testing.T) {
	type args struct {
		stackName string
		stackFile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "valid", args: args{
			stackName: "awsgrips-test-log-group",
			stackFile: "/Users/nmarks/go/src/github.com/natemarks/awsgrips/assets/cfntemplates/testloggroup.json",
		}, wantErr: false,
			want: "sdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateStack(tt.args.stackName, tt.args.stackFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.args.stackName) {
				t.Errorf("StackId does not contain stackName")
			}
			err = CreateStackWait(tt.args.stackName, 10)
			if err != nil {
				t.Errorf("Error waiting for stack to complete: %s", tt.args.stackName)
			}
			err = DeleteStack(tt.args.stackName)
			if err != nil {
				t.Errorf("unable to delete stack: %s", tt.args.stackName)
			}
		})
	}
}
