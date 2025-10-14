package main

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		{
			s: "",
			t: "",
			want: true,
		},
		{
			s: "a",
			t: "a",
			want: true,
		},
		{
			s: "ab",
			t: "a",
			want: false,
		},
		{
			s: "a",
			t: "ab",
			want: false,
		},
		{
			s: "cat",
			t: "tac",
			want: true,
		},
		{
			s: "t",
			t: "c",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("%v, %v", got, fmt.Sprintf("%v : %v", tt.s, tt.t))
			}
		})
	}
}

