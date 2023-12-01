package FindLuckyNumber

import (
	"testing"
)

func TestFindLucky(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test case 1", []int{1, 2, 2, 3, 3, 3}, 3},
		{"Test case 2", []int{1, 1, 2, 2, 3}, 2},
		{"Test case 3", []int{1, 1, 2, 2, 3, 3, 3}, 3},
		{"Test case 4", []int{1, 1, 1, 1, 1, 1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.input); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
