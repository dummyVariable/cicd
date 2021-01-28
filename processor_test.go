package main

import (
	"testing"
)

func Test_build(t *testing.T) {

	tests := []struct {
		name       string
		statements []string
		wantErr    bool
	}{
		{
			name:       "build single commands",
			statements: []string{"echo lol"},
			wantErr:    false,
		}, {
			name:       "build multiple commands",
			statements: []string{"echo lol", "echo lol"},
			wantErr:    false,
		}, {
			name:       "build no commands",
			statements: []string{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := build(tt.statements); (err != nil) != tt.wantErr {
				t.Errorf("build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
