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

func TestFirstAncestorDir(t *testing.T) {
	type args struct {
		workingDir string
		targetDir  string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{name: "single target with first and last seps", args: args{
			workingDir: "/aaa/bbb/ccc/my_target/ddd/fff/",
			targetDir:  "my_target",
		},
			wantResult: "/aaa/bbb/ccc/my_target/",
			wantErr:    false},
		{name: "single target with first sep", args: args{
			workingDir: "/aaa/bbb/ccc/my_target/ddd/fff",
			targetDir:  "my_target",
		},
			wantResult: "/aaa/bbb/ccc/my_target/",
			wantErr:    false},
		{name: "single target with netiher first last seps", args: args{
			workingDir: "aaa/bbb/ccc/my_target/ddd/fff",
			targetDir:  "my_target",
		},
			wantResult: "/aaa/bbb/ccc/my_target/",
			wantErr:    false},
		{name: "double target", args: args{
			workingDir: "/aaa/bbb/ccc/my_target/ddd/my_target/ggg/hhh",
			targetDir:  "my_target",
		},
			wantResult: "/aaa/bbb/ccc/my_target/ddd/my_target/",
			wantErr:    false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := FirstAncestorDir(tt.args.workingDir, tt.args.targetDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstAncestorDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("FirstAncestorDir() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
