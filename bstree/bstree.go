package bstree

import (
	"golang.org/x/exp/constraints"
)

// BSTree is a binary tree data structure.
type BSTree[V constraints.Ordered] struct {
	root *Node[V]
}

// New returns a new BinarySearchTree with the given value as the root.
func New[V constraints.Ordered](rootVal V) *BSTree[V] {
	return &BSTree[V]{
		root: NewNode(rootVal),
	}
}

// Root returns the root of the tree
func (bst *BSTree[V]) Root() *Node[V] {
	return bst.root
}

// Search returns the node with the given value if it exists in the tree.
func (bst *BSTree[V]) Search(root *Node[V], query V) *Node[V] {
	if root == nil || root.value == query {
		return root
	}

	if query < root.value {
		return bst.Search(root.left, query)
	}

	return bst.Search(root.right, query)
}

// Insert inserts a new node into the tree.
func (bst *BSTree[V]) Insert(root *Node[V], value V) *Node[V] {
	// if the root is nil, then we know we are inserting at this node
	if root == nil {
		return NewNode(value)
	}

	if value < root.value {
		// insert the node to the left of the root
		root.left = bst.Insert(root.left, value)
	} else if value > root.value {
		// insert the node to the right of the root
		root.right = bst.Insert(root.right, value)
	}

	return root
}

func (bst *BSTree[V]) InorderSuccessor(root *Node[V]) *Node[V] {
	curr := root
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

// Remove removes the node with the given value from the tree.
func (bst *BSTree[V]) Remove(node *Node[V], query V) *Node[V] {
	if node == nil {
		return nil
	}

	// Go left
	if query < node.value {
		node.left = bst.Remove(node.left, query)
		return node
	}
	if query > node.value {
		node.right = bst.Remove(node.right, query)
		return node
	}

	// node.value == query so we need to remove it
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	// If the node has only a right child, replace the node with its right child
	if node.left == nil {
		node = node.right
		return node
	}

	// If the node has only a left child, replace the node with its left child
	if node.right == nil {
		node = node.left
		return node
	}

	// If the node has two children, replace the node with its inorder successor
	inSuccessor := bst.InorderSuccessor(node.right)

	// Copy the value from the inorder successor to the node
	node.value = inSuccessor.value

	// Delete the inorder successor now that it's value has been copied
	node.right = bst.Remove(node.right, inSuccessor.value)

	return node
}

// Inorder returns the values of the tree in order.
func (bst *BSTree[V]) Inorder(root *Node[V]) []V {
	nodes := make([]V, 0)
	bst.inorder(root, &nodes)
	return nodes

}

func (bst *BSTree[V]) inorder(root *Node[V], dst *[]V) {
	if root != nil {
		bst.inorder(root.left, dst)
		*dst = append(*dst, root.value)
		bst.inorder(root.right, dst)
	}
}

// Preorder returns the values of the tree in preorder.
func (bst *BSTree[V]) Preorder(root *Node[V]) []V {
	nodes := make([]V, 0)
	bst.preorder(root, &nodes)
	return nodes

}

func (bst *BSTree[V]) preorder(root *Node[V], dst *[]V) {
	if root != nil {
		*dst = append(*dst, root.value)
		bst.preorder(root.left, dst)
		bst.preorder(root.right, dst)
	}
}

// Postorder returns the values of the tree in postorder.
func (bst *BSTree[V]) Postorder(root *Node[V]) []V {
	nodes := make([]V, 0)
	bst.postorder(root, &nodes)
	return nodes

}

func (bst *BSTree[V]) postorder(root *Node[V], dst *[]V) {
	if root != nil {
		bst.postorder(root.left, dst)
		bst.postorder(root.right, dst)
		*dst = append(*dst, root.value)
	}
}
