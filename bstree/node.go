package bstree

// Node is a node in a binary search tree.
type Node[T any] struct {
	value       T
	left, right *Node[T]
}

// NewNode returns a new node with the given value.
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

// Value returns the value of the node.
func (n *Node[T]) Value() T {
	return n.value
}

// Left returns the left child of the node.
func (n *Node[T]) Left() *Node[T] {
	return n.left
}

// Right returns the right child of the node.
func (n *Node[T]) Right() *Node[T] {
	return n.right
}
