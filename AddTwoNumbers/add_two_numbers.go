package AddTwoNumbers

import (
	"strconv"
)

/*
  Definition for singly-linked list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers is a function that adds two numbers represented as linked lists.
// It returns a new linked list that represents the sum of the two input linked lists.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var prev *ListNode
	var result *ListNode

	// Loop through the linked lists until both l1 and l2 are nil.
	for l1 != nil || l2 != nil {
		sum := carry

		// If l1 is not nil, add l1.Val to sum and move l1 to the next node.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		// If l2 is not nil, add l2.Val to sum and move l2 to the next node.
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		// Create a new node with the value of sum%10 and add it to the result linked list.
		tempNode := &ListNode{Val: sum % 10}
		if prev == nil {
			result = tempNode
			prev = tempNode
		} else {
			prev.Next = tempNode
			prev = prev.Next
		}
	}
	// If there is a carry after the loop, add it to the result linked list.
	if carry > 0 {
		tempNode := &ListNode{Val: carry}
		prev.Next = tempNode
	}
	return result
}

// LinkedListToString is a method for a ListNode struct. It traverses a linked list and
// converts it into a string representation in the format "7 -> 8 -> 9".
func (node *ListNode) LinkedListToString() string {
	var result string
	/*
		Create a temporary pointer to traverse the linked list without modifying the original node.
		This eliminates the warning about Assignment to the method receiver propagates only to calls but not to callers.
	*/
	temp := node
	for temp != nil {
		result += strconv.Itoa(temp.Val) + " -> "
		temp = temp.Next
	}
	// Remove the last " -> "
	if len(result) > 4 {
		result = result[:len(result)-4]
	}
	return result
}
