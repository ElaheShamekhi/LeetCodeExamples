package MajorityElement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test case 1", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"Test case 2", []int{1, 1, 2, 2, 2, 2, 2, 3, 3}, 2},
		{"Test case 3", []int{1, 2, 3, 4, 5, 6}, -1},
		{"Test case 4", []int{1, 1, 1, 2, 2, 2}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.input); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
