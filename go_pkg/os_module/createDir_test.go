package main

import (
	"testing"
)

func TestMkdir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "make dir test",
			args: args{"./mydir"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mkdir(tt.args.path)
		})
	}
}

// func TestExpandDir(t *testing.T) {
// 	type args struct {
// 		dir string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{name: "absulte path test",
// 			args: args{dir: "./mydir"},
// 			want: "mydir"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ExpandDir(tt.args.dir); got != tt.want {
// 				t.Errorf("ExpandDir() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
