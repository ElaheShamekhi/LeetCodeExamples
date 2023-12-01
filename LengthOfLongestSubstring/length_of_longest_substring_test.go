package LengthOfLongestSubstring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Test case 1", "abcabcbb", 3},
		{"Test case 2", "bbbbb", 1},
		{"Test case 3", "pwwkew", 3},
		{"Test case 4", "abcdef", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.input); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
