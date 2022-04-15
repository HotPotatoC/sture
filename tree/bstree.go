package tree

// BSTree is a binary tree data structure.
type BSTree[K, V any] struct {
	root *BSTNode[K, V]
	cmp  func(K, K) int
}

// NewBinarySearchTree returns a new binary search tree with the given value as the root.
// The second parameter is a function that returns -1, 0, or 1 depending on whether x < y, x == y, or x > y
// Example Usage:
//	bstree.New(1, sture.Compare[int])
func NewBinarySearchTree[K, V any](cmp func(a, b K) int) *BSTree[K, V] {
	return &BSTree[K, V]{
		cmp: cmp,
	}
}

// Root returns the root of the tree
func (bst *BSTree[K, V]) Root() *BSTNode[K, V] {
	return bst.root
}

// Search returns the node with the given value if it exists in the tree.
func (bst *BSTree[K, V]) Search(query K) *BSTNode[K, V] {
	if bst.root == nil || bst.cmp(bst.root.key, query) == 0 {
		return bst.root
	}

	curr := bst.root
	for curr != nil {
		if bst.cmp(curr.key, query) == 0 {
			return curr
		} else if bst.cmp(curr.key, query) > 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return nil
}

// insert inserts a new node into the tree.
func (bst *BSTree[K, V]) insert(node *BSTNode[K, V], key K, value V) *BSTNode[K, V] {
	// if the node is nil, then we know we are inserting at this node
	if node == nil {
		return NewBSTNode(key, value)
	}

	if bst.cmp(key, node.key) < 0 {
		node.left = bst.insert(node.left, key, value)
	} else if bst.cmp(key, node.key) > 0 {
		node.right = bst.insert(node.right, key, value)
	} else { // key == node.key
		node.value = value // update value
	}

	return node
}

// remove removes the node with the given value from the tree. (internal-use only)
func (bst *BSTree[K, V]) remove(node *BSTNode[K, V], key K) *BSTNode[K, V] {
	if node == nil {
		return nil
	}

	if bst.cmp(key, node.key) < 0 {
		node.left = bst.remove(node.left, key)
	} else if bst.cmp(key, node.key) > 0 {
		node.right = bst.remove(node.right, key)
	} else { // key == node.key
		if node.left == nil && node.right == nil { // no children
			return nil
		} else if node.left == nil { // one child
			return node.right
		} else if node.right == nil {
			return node.left
		} else { // two children
			// find the in-order successor
			successor := bst.InorderSuccessor(node.right)
			node.key = successor.key
			node.value = successor.value
			node.right = bst.remove(node.right, successor.key)
		}
	}

	return node
}

// Insert inserts a new node into the tree.
func (bst *BSTree[K, V]) Insert(key K, value V) {
	bst.root = bst.insert(bst.root, key, value)
}

// InorderSuccessor returns the inorder successor of the given node.
func (bst *BSTree[K, V]) InorderSuccessor(root *BSTNode[K, V]) *BSTNode[K, V] {
	curr := root
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

// Remove removes the node with the given value from the tree.
func (bst *BSTree[K, V]) Remove(key K) {
	bst.remove(bst.root, key)
}

// Inorder returns the values of the tree in order.
func (bst *BSTree[K, V]) Inorder() []*BSTNode[K, V] {
	nodes := make([]*BSTNode[K, V], 0)
	bst.inorder(bst.root, &nodes)
	return nodes

}

func (bst *BSTree[K, V]) inorder(node *BSTNode[K, V], dst *[]*BSTNode[K, V]) {
	if node != nil {
		bst.inorder(node.left, dst)
		*dst = append(*dst, node)
		bst.inorder(node.right, dst)
	}
}

// Preorder returns the values of the tree in preorder.
func (bst *BSTree[K, V]) Preorder() []*BSTNode[K, V] {
	nodes := make([]*BSTNode[K, V], 0)
	bst.preorder(bst.root, &nodes)
	return nodes

}

func (bst *BSTree[K, V]) preorder(node *BSTNode[K, V], dst *[]*BSTNode[K, V]) {
	if node != nil {
		*dst = append(*dst, node)
		bst.preorder(node.left, dst)
		bst.preorder(node.right, dst)
	}
}

// Postorder returns the values of the tree in postorder.
func (bst *BSTree[K, V]) Postorder() []*BSTNode[K, V] {
	nodes := make([]*BSTNode[K, V], 0)
	bst.postorder(bst.root, &nodes)
	return nodes

}

func (bst *BSTree[K, V]) postorder(root *BSTNode[K, V], dst *[]*BSTNode[K, V]) {
	if root != nil {
		bst.postorder(root.left, dst)
		bst.postorder(root.right, dst)
		*dst = append(*dst, root)
	}
}
