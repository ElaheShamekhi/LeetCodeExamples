package LongestAscendingSubarray

import (
	"fmt"
	"testing"
)

type test struct {
	input []int
	want  int
}

func TestLongestAscendingSubarray(t *testing.T) {
	tests := []test{
		{input: []int{5, 6, 3, 5, 7, 8, 9, 1, 2}, want: 5},
		{input: []int{1, 2, 3, 4, 5}, want: 5},
		{input: []int{5, 4, 3, 2, 1}, want: 1},
		{input: []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}, want: 5},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test number %d", i+1)
		t.Run(testName, func(t *testing.T) {
			got := longestAscendingSubarray(tt.input)
			if got != tt.want {
				t.Errorf("longestAscendingSubarray(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
