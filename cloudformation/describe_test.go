//go:build !unit
// +build !unit

package cloudformation

import (
	"testing"
)

func TestGetStackByNameSubstring1(t *testing.T) {
	type args struct {
		sub string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid", args: args{sub: "idp"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStackByNameSubstring(tt.args.sub)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStackByNameSubstring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < 1 {
				t.Error("no matching stacks found")
			}
		})
	}
}
