package gotest

import "testing"

func Test_addadd(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addadd()
		})
	}
}
