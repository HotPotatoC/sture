package bstree

// BSTree is a binary tree data structure.
type BSTree[V any] struct {
	root *Node[V]
	cmp  func(V, V) int
}

// New returns a new binary search tree with the given value as the root.
// The second parameter is a function that returns -1, 0, or 1 depending on whether x < y, x == y, or x > y
// Example Usage:
//	bstree.New(1, sture.Compare[int])
func New[V any](rootVal V, cmp func(a, b V) int) *BSTree[V] {
	return &BSTree[V]{
		root: NewNode(rootVal),
		cmp:  cmp,
	}
}

// Root returns the root of the tree
func (bst *BSTree[V]) Root() *Node[V] {
	return bst.root
}

// Search returns the node with the given value if it exists in the tree.
func (bst *BSTree[V]) Search(root *Node[V], query V) *Node[V] {
	if root == nil || bst.cmp(root.value, query) == 0 {
		return root
	}

	for root != nil {
		if bst.cmp(root.value, query) == 0 {
			return root
		} else if bst.cmp(root.value, query) > 0 {
			root = root.left
		} else {
			root = root.right
		}
	}

	return nil
}

// Insert inserts a new node into the tree.
func (bst *BSTree[V]) Insert(root *Node[V], value V) *Node[V] {
	// if the root is nil, then we know we are inserting at this node
	if root == nil {
		return NewNode(value)
	}

	for root != nil {
		if bst.cmp(root.value, value) == 0 {
			// the value is already in the tree, so we don't need to insert it
			return root
		} else if bst.cmp(root.value, value) > 0 {
			// if the left child is nil, then we know we are inserting at this node
			if root.left == nil {
				root.left = NewNode(value)
				return root.left
			}
			// otherwise we keep traversing down the left side of the tree
			root = root.left
		} else {
			// if the right child is nil, then we know we are inserting at this node
			if root.right == nil {
				root.right = NewNode(value)
				return root.right
			}
			// otherwise we keep traversing down the right side of the tree
			root = root.right
		}
	}

	return nil
}

// InorderSuccessor returns the inorder successor of the given node.
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

	if bst.cmp(node.value, query) > 0 {
		node.left = bst.Remove(node.left, query)
		return node
	}

	if bst.cmp(node.value, query) < 0 {
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
