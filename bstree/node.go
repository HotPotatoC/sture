package bstree

// Node is a node in a binary search tree.
type Node[V any] struct {
	value       V
	left, right *Node[V]
}

// NewNode returns a new node with the given value.
func NewNode[V any](value V) *Node[V] {
	return &Node[V]{value: value}
}

// Value returns the value of the node.
func (n *Node[V]) Value() V {
	return n.value
}

// Left returns the left child of the node.
func (n *Node[V]) Left() *Node[V] {
	return n.left
}

// Right returns the right child of the node.
func (n *Node[V]) Right() *Node[V] {
	return n.right
}
