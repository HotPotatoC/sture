package tree

// BSTNode is a node in a binary search tree.
type BSTNode[K, V any] struct {
	key         K
	value       V
	left, right *BSTNode[K, V]
}

// NewBSTNode returns a new node with the given value.
func NewBSTNode[K, V any](key K, value V) *BSTNode[K, V] {
	return &BSTNode[K, V]{key: key, value: value}
}

// Key returns the key of the node.
func (n *BSTNode[K, V]) Key() K {
	return n.key
}

// Value returns the value of the node.
func (n *BSTNode[K, V]) Value() V {
	return n.value
}

// Left returns the left child of the node.
func (n *BSTNode[K, V]) Left() *BSTNode[K, V] {
	return n.left
}

// Right returns the right child of the node.
func (n *BSTNode[K, V]) Right() *BSTNode[K, V] {
	return n.right
}
