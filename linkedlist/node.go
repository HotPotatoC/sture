package linkedlist

// Node is a node in a linked list.
type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// NewNode returns a new linked list node.
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

// Value returns the value of the node.
func (node *Node[T]) Value() T {
	return node.value
}

// Next returns the next node in the list.
func (node *Node[T]) Next() *Node[T] {
	return node.next
}

// Prev returns the previous node in the list.
func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}
