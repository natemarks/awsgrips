package testfunctions

import (
	"os"
	"testing"
)

func TestJSONFileToByteArray(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid",
			args: args{
				filePath: getProjectPath() + "test/data/someData.json",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd, _ := os.Getwd()
			t.Log(wd)
			_, err := JSONFileToByteArray(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONFileToByteArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
