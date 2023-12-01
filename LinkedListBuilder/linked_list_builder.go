package LinkedListBuilder

type ListNode struct {
	Val  int
	Next *ListNode
}

type Builder struct {
	head *ListNode
}

func NewBuilder() *Builder {
	return &Builder{}
}

// Add adds a new node with the given value to the linked list.
func (b *Builder) Add(val int) *Builder {
	newNode := &ListNode{Val: val}
	if b.head == nil {
		b.head = newNode
	} else {
		current := b.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	return b
}

// Delete deletes the first node in the linked list that has the given value.
func (b *Builder) Delete(val int) *Builder {
	if b.head == nil {
		return b
	}
	// If the head of the list has the given value, update the head to the next node.
	if b.head.Val == val {
		b.head = b.head.Next
		return b
	}
	// Initialize a pointer to the head of the list.
	current := b.head
	// Traverse the list until we find a node that has the next node with the given value.
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}
	// If we found such a node, update its next pointer to skip the next node.
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	return b
}

// Length calculates the length of the linked list.
func (b *Builder) Length() int {
	current := b.head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// Get returns the node at the given index in the linked list.
func (b *Builder) Get(index int) *ListNode {
	current := b.head
	for i := 0; i < index; i++ {
		if current != nil {
			current = current.Next
		} else {
			return nil
		}
	}
	return current
}

// GetJthFromEnd is a method of the Builder struct that returns the jth node from the end of the linked list.
func (b *Builder) GetJthFromEnd(j int) *ListNode {
	// Initialize a runner pointer at the head of the list.
	runner := b.head
	// Move the runner j steps from the head of the list.
	for i := 0; i < j; i++ {
		// If the runner reaches the end of the list before j steps, return nil.
		if runner == nil {
			return nil
		}
		// Move the runner one step forward.
		runner = runner.Next
	}
	// Initialize a current pointer at the head of the list.
	current := b.head
	// Move both the runner and current pointers one step at a time until the runner reaches the end of the list.
	for runner != nil {
		runner = runner.Next
		current = current.Next
	}
	// When the runner reaches the end of the list, the current pointer is at the jth element from the end of the list.
	// Return the current pointer.
	return current
}

// Build returns the head of the linked list.
func (b *Builder) Build() *ListNode {
	return b.head
}
