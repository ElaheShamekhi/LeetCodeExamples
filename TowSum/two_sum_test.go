package TowSum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestTable struct {
	Nums   []int
	Target int
	Want   []int
}

func InitializedTestTable() []TestTable {
	return []TestTable{
		{Nums: []int{3, 4, 6, 10, 15}, Target: 10, Want: []int{1, 2}},
		{Nums: []int{3, 6, 10, 15}, Target: 6, Want: []int{}},
		{Nums: []int{3, 3}, Target: 6, Want: []int{0, 1}},
		{Nums: []int{2, 7, 11, 15}, Target: 9, Want: []int{0, 1}},
		{Nums: []int{3, 2, 4}, Target: 6, Want: []int{1, 2}},
	}
}

func TestTwoSum(t *testing.T) {
	for i, tt := range InitializedTestTable() {
		testName := fmt.Sprintf("test number %d", i+1)
		t.Run(testName, func(t *testing.T) {
			ans := TwoSum(tt.Nums, tt.Target)
			if !reflect.DeepEqual(ans, tt.Want) {
				t.Errorf("got %v, want %v", ans, tt.Want)
			}
		})
	}

}
