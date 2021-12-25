package cloudformation

import "testing"

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
			stackName: "awsgrips_test_log_group",
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
			if got != tt.want {
				t.Errorf("CreateStack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
