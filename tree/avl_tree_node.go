package tree

// AVLNode is a node in a binary search tree.
type AVLNode[K any, V any] struct {
	key         K
	value       V
	height      int
	left, right *AVLNode[K, V]
}

// NewAVLNode returns a new node with the given value.
func NewAVLNode[K any, V any](key K, value V) *AVLNode[K, V] {
	return &AVLNode[K, V]{
		key:    key,
		value:  value,
		height: 1,
	}
}

// Key returns the key of the node.
func (n *AVLNode[K, V]) Key() K {
	return n.key
}

// Value returns the value of the node.
func (n *AVLNode[K, V]) Value() V {
	return n.value
}

// Height returns the height of the node.
func (n *AVLNode[K, V]) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

// Left returns the left child of the node.
func (n *AVLNode[K, V]) Left() *AVLNode[K, V] {
	return n.left
}

// Right returns the right child of the node.
func (n *AVLNode[K, V]) Right() *AVLNode[K, V] {
	return n.right
}
