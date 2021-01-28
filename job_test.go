package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_jobs(t *testing.T) {
	test1, _ := ioutil.ReadFile("test1.yaml")
	test2, _ := ioutil.ReadFile("test2.yaml")

	tests := []struct {
		name string
		args []byte
		want Job
	}{
		{
			name: "test for single command",
			args: test1,
			want: Job{
				PreBuild:  []string{"echo prebuild stage"},
				Build:     []string{"echo build stage"},
				PostBuild: []string{"echo postbuild stage"},
			},
		},
		{
			name: "test for multiple command",
			args: test2,
			want: Job{
				PreBuild:  []string{"echo prebuild stage", "echo extra step 1"},
				Build:     []string{"echo build stage", "echo extra step 2"},
				PostBuild: []string{"echo postbuild stage", "echo extra step 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jobs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
