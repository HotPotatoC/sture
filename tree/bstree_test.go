package tree_test

import (
	"reflect"
	"testing"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/tree"
)

func TestBStree_Insert(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(8, 8)
	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)

	if bst.Root().Value() != 8 {
		t.Errorf("bst.Root().Value() = %v, want %v", bst.Root().Value(), 8)
	}

	if bst.Root().Left().Value() != 3 {
		t.Errorf("bst.Root().Left().Value() = %v, want %v", bst.Root().Left().Value(), 3)
	}

	if bst.Root().Right().Value() != 10 {
		t.Errorf("bst.Root().Right().Value() = %v, want %v", bst.Root().Right().Value(), 10)
	}

	if bst.Root().Left().Left().Value() != 1 {
		t.Errorf("bst.Root().Left().Left().Value() = %v, want %v", bst.Root().Left().Left().Value(), 1)
	}

	if bst.Root().Left().Right().Value() != 6 {
		t.Errorf("bst.Root().Left().Right().Value() = %v, want %v", bst.Root().Left().Right().Value(), 6)
	}

	if bst.Root().Left().Right().Left().Value() != 4 {
		t.Errorf("bst.Root().Left().Right().Left().Value() = %v, want %v", bst.Root().Left().Right().Left().Value(), 4)
	}
}

func TestBSTree_Search(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)

	if bst.Search(3) == nil {
		t.Errorf("bst.Search(3) = nil, want not nil")
	}

	if bst.Search(10) == nil {
		t.Errorf("bst.Search(10) = nil, want not nil")
	}

	if bst.Search(1) == nil {
		t.Errorf("bst.Search(1) = nil, want not nil")
	}

	if bst.Search(6) == nil {
		t.Errorf("bst.Search(6) = nil, want not nil")
	}

	if bst.Search(4) == nil {
		t.Errorf("bst.Search(4) = nil, want not nil")
	}

	if bst.Search(9) != nil {
		t.Errorf("bst.Search(9) = %v, want nil", bst.Search(9))
	}
}

func TestBSTree_Remove(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)
	bst.Insert(7, 7)

	bst.Remove(3)

	if bst.Root().Value() != 4 {
		t.Errorf("bst.Root().Value() = %v, want %v", bst.Root().Value(), 4)
	}

	if bst.Root().Left().Value() != 1 {
		t.Errorf("bst.Root().Left().Value() = %v, want %v", bst.Root().Left().Value(), 1)
	}

	if bst.Root().Right().Value() != 10 {
		t.Errorf("bst.Root().Right().Value() = %v, want %v", bst.Root().Right().Value(), 10)
	}
}

func TestBSTree_Inorder(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)

	list := bst.Inorder()
	exp := []int{1, 3, 4, 6, 10}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Inorder() = %v, want %v", list, exp)
	}
}

func TestBSTree_Preorder(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)

	list := bst.Preorder()
	exp := []int{3, 1, 10, 6, 4}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Preorder() = %v, want %v", list, exp)
	}
}

func TestBSTree_Postorder(t *testing.T) {
	bst := tree.NewBinarySearchTree[int, int](sture.Compare[int])

	bst.Insert(3, 3)
	bst.Insert(10, 10)
	bst.Insert(1, 1)
	bst.Insert(6, 6)
	bst.Insert(4, 4)

	list := bst.Postorder()
	exp := []int{1, 4, 6, 10, 3}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Postorder() = %v, want %v", list, exp)
	}
}
