package linkedlist

// Node is a node in a linked list.
type Node[V any] struct {
	value V
	next  *Node[V]
	prev  *Node[V]
}

// NewNode returns a new linked list node.
func NewNode[V any](value V) *Node[V] {
	return &Node[V]{value: value}
}

// Value returns the value of the node.
func (node *Node[V]) Value() V {
	return node.value
}

// Next returns the next node in the list.
func (node *Node[V]) Next() *Node[V] {
	return node.next
}

// Prev returns the previous node in the list.
func (node *Node[V]) Prev() *Node[V] {
	return node.prev
}
