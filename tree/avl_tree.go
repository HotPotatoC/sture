package tree

import (
	"github.com/HotPotatoC/sture"
)

// AVLTree is a self-balancing binary search tree.
type AVLTree[K any, V any] struct {
	// root is the root of the tree
	root *AVLNode[K, V]
	// cmp is a function that returns -1, 0, or 1 depending on whether x < y, x == y, or x > y
	// (used for comparing keys)
	cmp func(K, K) int
}

// New returns a new avl tree. The first parameter is a function that returns -1, 0, or 1 depending on whether x < y, x == y, or x > y
// Example Usage:
//	tree.NewAVLTree(sture.Compare[T]) // where T is any type that supports the operators < <= >= >
//	tree.NewAVLTree(bytes.Compare) // for byte slices
func NewAVLTree[K any, V any](cmp func(a K, b K) int) *AVLTree[K, V] {
	return &AVLTree[K, V]{
		cmp: cmp,
	}
}

// calcBalanceFactor calculates the balance factor of the given node.
// BF(T) = h(LT) - h(RT) = {-1, 0, 1}, where h(LT) =  height of Left Sub-tree
// and h(RT) = Right Sub-tree. If the balance factor of the node is greater than 1 or
// less than -1 therefore the node is unbalanced.
func (avl *AVLTree[K, V]) calcBalanceFactor(node *AVLNode[K, V]) int {
	if node == nil {
		return 0
	}

	return node.left.Height() - node.right.Height()
}

// rebalance rebalances the tree at the given node.
func (avl *AVLTree[K, V]) rebalance(node *AVLNode[K, V]) *AVLNode[K, V] {
	node.height = 1 + sture.Max(node.left.Height(), node.right.Height()) // update height
	balanceFactor := avl.calcBalanceFactor(node)                         // get balance factor

	// balanceFactor > 1 means the left sub tree is heavy
	if balanceFactor > 1 {
		if avl.calcBalanceFactor(node.left) < 0 { // left left heavy
			node.left = avl.RotateLeft(node.left) // double rotation
		}

		return avl.RotateRight(node)
	}

	// if the tree is right heavy
	if balanceFactor < -1 {
		if avl.calcBalanceFactor(node.right) > 0 { // right right heavy
			node.right = avl.RotateRight(node.right) // double rotation
		}

		return avl.RotateLeft(node)
	}

	return node
}

// insert inserts a new node into the tree and attempts to rebalance the tree (internal-use only).
func (avl *AVLTree[K, V]) insert(node *AVLNode[K, V], key K, value V) *AVLNode[K, V] {
	// if the node is nil, then we know we are inserting at this node
	if node == nil {
		return NewAVLNode(key, value)
	}

	if avl.cmp(key, node.key) < 0 {
		node.left = avl.insert(node.left, key, value)
	} else if avl.cmp(key, node.key) > 0 {
		node.right = avl.insert(node.right, key, value)
	} else { // key == node.key
		node.value = value // update value
	}

	return avl.rebalance(node)
}

// remove removes the node with the given value from the tree and attempts to rebalance the tree (internal-use only).
func (avl *AVLTree[K, V]) remove(node *AVLNode[K, V], query K) *AVLNode[K, V] {
	if node == nil {
		return nil
	}

	if avl.cmp(query, node.key) < 0 {
		node.left = avl.remove(node.left, query)
	} else if avl.cmp(query, node.key) > 0 {
		node.right = avl.remove(node.right, query)
	} else { // key == node.key
		if node.left == nil && node.right == nil { // no children
			return nil
		} else if node.left == nil { // one child
			return node.right
		} else if node.right == nil {
			return node.left
		} else { // two children
			// find the in-order successor
			successor := avl.InorderSuccessor(node.right)
			node.key = successor.key
			node.value = successor.value
			node.right = avl.remove(node.right, successor.key)
		}
	}

	return avl.rebalance(node)
}

// Root returns the root of the tree
func (avl *AVLTree[K, V]) Root() *AVLNode[K, V] {
	return avl.root
}

// InorderSuccessor returns the inorder successor of the given node.
func (avl *AVLTree[K, V]) InorderSuccessor(node *AVLNode[K, V]) *AVLNode[K, V] {
	curr := node
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

// RotateLeft rotates the tree to the left.
func (avl *AVLTree[K, V]) RotateLeft(node *AVLNode[K, V]) *AVLNode[K, V] {
	pivot := node.right
	pivotLChild := pivot.left
	pivot.left = node
	node.right = pivotLChild
	node.height = sture.Max(node.left.Height(), node.right.Height()) + 1
	pivot.height = sture.Max(pivot.left.Height(), pivot.right.Height()) + 1
	return pivot
}

// RotateRight rotates the tree to the right.
func (avl *AVLTree[K, V]) RotateRight(node *AVLNode[K, V]) *AVLNode[K, V] {
	pivot := node.left
	pivotRChild := pivot.right
	pivot.right = node
	node.left = pivotRChild
	pivot.height = sture.Max(pivot.left.Height(), pivot.right.Height()) + 1
	node.height = sture.Max(node.left.Height(), node.right.Height()) + 1
	return pivot
}

// Insert inserts a new node into the tree and attempts to rebalance the tree.
func (avl *AVLTree[K, V]) Insert(key K, value V) {
	avl.root = avl.insert(avl.root, key, value)
}

// Search searches for the given key in the tree and returns the value if found.
func (avl *AVLTree[K, V]) Search(key K) *AVLNode[K, V] {
	// if the tree is empty, return zero-value of V or
	// if the root's key equals to the given search query
	// return it immediately
	if avl.root == nil || avl.cmp(key, avl.root.key) == 0 {
		return avl.root
	}

	node := avl.root
	for node != nil {
		// if the node's key equals to the given search query
		if avl.cmp(key, node.key) == 0 {
			return node
		} else if avl.cmp(key, node.key) < 0 {
			node = node.left // go left
		} else {
			node = node.right // go right
		}
	}

	return nil
}

// Remove removes the node with the given key from the tree and attempts to rebalance the tree.
func (avl *AVLTree[K, V]) Remove(query K) {
	avl.root = avl.remove(avl.root, query)
}

// Inorder returns the values of the tree in order.
func (avl *AVLTree[K, V]) Inorder(root *AVLNode[K, V]) []V {
	nodes := make([]V, 0)
	avl.inorder(root, &nodes)
	return nodes

}

func (avl *AVLTree[K, V]) inorder(root *AVLNode[K, V], dst *[]V) {
	if root != nil {
		avl.inorder(root.left, dst)
		*dst = append(*dst, root.value)
		avl.inorder(root.right, dst)
	}
}

// Preorder returns the values of the tree in preorder.
func (avl *AVLTree[K, V]) Preorder(root *AVLNode[K, V]) []V {
	nodes := make([]V, 0)
	avl.preorder(root, &nodes)
	return nodes

}

func (avl *AVLTree[K, V]) preorder(root *AVLNode[K, V], dst *[]V) {
	if root != nil {
		*dst = append(*dst, root.value)
		avl.preorder(root.left, dst)
		avl.preorder(root.right, dst)
	}
}

// Postorder returns the values of the tree in postorder.
func (avl *AVLTree[K, V]) Postorder(root *AVLNode[K, V]) []V {
	nodes := make([]V, 0)
	avl.postorder(root, &nodes)
	return nodes

}

func (avl *AVLTree[K, V]) postorder(root *AVLNode[K, V], dst *[]V) {
	if root != nil {
		avl.postorder(root.left, dst)
		avl.postorder(root.right, dst)
		*dst = append(*dst, root.value)
	}
}
