package logger

import (
	"reflect"
	"testing"
)

func Test_customCallerMarshalFunc(t *testing.T) {
	filepath := "/home/xyx/code/suite-gt/common/customization_test.go"
	type args struct {
		basepath string
		fullpath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return custom caller",
			args: args{
				basepath: "common",
				fullpath: filepath,
			},
			want: "common/customization_test.go:35",
		},
		{
			name: "return the full path",
			args: args{
				basepath: "",
				fullpath: filepath,
			},
			want: "/home/xyx/code/suite-gt/common/customization_test.go:35",
		},
		{
			name: "return file only",
			args: args{
				basepath: "typo",
				fullpath: filepath,
			},
			want: "customization_test.go:35",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := customCallerMarshalFunc(tt.args.basepath)
			got := res(uintptr(0), tt.args.fullpath, 35)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customCallerMarshalFunc()"+
					"\ngot:  %v"+
					"\nwant: %v", got, tt.want)
			}
		})
	}
}
