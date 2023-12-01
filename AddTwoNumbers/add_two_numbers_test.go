package AddTwoNumbers

import (
	"fmt"
	"reflect"
	"testing"
)

type TestTable struct {
	FirstList  *ListNode
	SecondList *ListNode
	Want       *ListNode
}

func InitializedTestTable() []TestTable {
	return []TestTable{
		{
			FirstList:  &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			SecondList: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			Want:       &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			FirstList:  &ListNode{Val: 0},
			SecondList: &ListNode{Val: 0},
			Want:       &ListNode{Val: 0},
		},
		{
			FirstList:  &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			SecondList: &ListNode{Val: 9},
			Want:       &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}},
		},
	}
}

func TestAddTwoNumbers(t *testing.T) {
	for i, tt := range InitializedTestTable() {
		testName := fmt.Sprintf("test number %d", i+1)
		t.Run(testName, func(t *testing.T) {
			ans := AddTwoNumbers(tt.FirstList, tt.SecondList)
			if !reflect.DeepEqual(ans, tt.Want) {
				t.Errorf("got %v, want %v", ans, tt.Want)
			}
		})
	}

}
