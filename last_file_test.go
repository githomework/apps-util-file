package file_test

import (
	"testing"

	"github.com/githomework/apps-app"
	"github.com/githomework/apps-util-file"
)

func TestWithPrefix(t *testing.T) {
	var a app.AppType
	a.Setup()
	type args struct {
		dir    string
		prefix string
	}
	folder := a.Folder
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{
		name:    "1",
		args:    args{dir: folder, prefix: "x"},
		want:    "x",
		wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := file.LastWithPrefix(tt.args.dir, tt.args.prefix)

			t.Log("**** Please manual check:", got, folder)

		})
	}
}
