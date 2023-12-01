package LinkedListBuilder

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	b := NewBuilder()
	b.Add(1).Add(2).Add(3)
	if b.Length() != 3 {
		t.Fatalf("Expected length 3, got %d", b.Length())
	}
}

func TestDelete(t *testing.T) {
	b := NewBuilder()
	b.Add(1).Add(2).Add(3)
	b.Delete(2)
	if b.Length() != 2 {
		t.Fatalf("Expected length 2, got %d", b.Length())
	}
}

func TestLength(t *testing.T) {
	b := NewBuilder()
	b.Add(1).Add(2).Add(3).Add(4)
	if b.Length() != 4 {
		t.Fatalf("Expected length 3, got %d", b.Length())
	}
}

func TestGet(t *testing.T) {
	b := NewBuilder()
	b.Add(1).Add(2).Add(3)
	node := b.Get(1)
	if node == nil || node.Val != 2 {
		t.Fatalf("Expected node with value 2, got %v", node)
	}
}

func TestGetJthFromEnd(t *testing.T) {
	b := NewBuilder()
	b.Add(1).Add(2).Add(3).Add(4).Add(5)

	// Test cases for jth element from end
	tests := []struct {
		j    int
		want int
	}{
		{j: 1, want: 5},
		{j: 2, want: 4},
		{j: 3, want: 3},
		{j: 4, want: 2},
		{j: 5, want: 1},
	}

	// Run the test cases
	for i, tt := range tests {
		testName := fmt.Sprintf("test number %d", i+1)
		t.Run(testName, func(t *testing.T) {
			node := b.GetJthFromEnd(tt.j)
			if node == nil || node.Val != tt.want {
				t.Errorf("GetJthFromEnd(%d) = %v, want %v", tt.j, node, tt.want)
			}
		})
	}
}
