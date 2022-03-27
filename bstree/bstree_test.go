package bstree_test

import (
	"reflect"
	"testing"

	"github.com/HotPotatoC/sture/bstree"
)

func TestBStree_Insert(t *testing.T) {
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)

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
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)

	if bst.Search(bst.Root(), 3) == nil {
		t.Errorf("bst.Search(bst.Root(), 3) = nil, want not nil")
	}

	if bst.Search(bst.Root(), 10) == nil {
		t.Errorf("bst.Search(bst.Root(), 10) = nil, want not nil")
	}

	if bst.Search(bst.Root(), 1) == nil {
		t.Errorf("bst.Search(bst.Root(), 1) = nil, want not nil")
	}

	if bst.Search(bst.Root(), 6) == nil {
		t.Errorf("bst.Search(bst.Root(), 6) = nil, want not nil")
	}

	if bst.Search(bst.Root(), 4) == nil {
		t.Errorf("bst.Search(bst.Root(), 4) = nil, want not nil")
	}

	if bst.Search(bst.Root(), 9) != nil {
		t.Errorf("bst.Search(bst.Root(), 9) = %v, want nil", bst.Search(bst.Root(), 9))
	}
}

func TestBSTree_Remove(t *testing.T) {
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)
	bst.Insert(bst.Root(), 7)

	bst.Remove(bst.Root(), 3)

	if bst.Root().Value() != 8 {
		t.Errorf("bst.Root().Value() = %v, want %v", bst.Root().Value(), 8)
	}

	if bst.Root().Left().Value() != 4 {
		t.Errorf("bst.Root().Left().Value() = %v, want %v", bst.Root().Left().Value(), 4)
	}

	bst.Remove(bst.Root(), 6)

	if bst.Root().Left().Right().Value() != 7 {
		t.Errorf("bst.Root().Left().Right().Value() = %v, want %v", bst.Root().Left().Right().Value(), 7)
	}
}

func TestBSTree_Inorder(t *testing.T) {
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)

	list := bst.Inorder(bst.Root())
	exp := []int{1, 3, 4, 6, 8, 10}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Inorder(bst.Root()) = %v, want %v", list, exp)
	}
}

func TestBSTree_Preorder(t *testing.T) {
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)

	list := bst.Preorder(bst.Root())
	exp := []int{8, 3, 1, 6, 4, 10}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Preorder(bst.Root()) = %v, want %v", list, exp)
	}
}

func TestBSTree_Postorder(t *testing.T) {
	bst := bstree.New(8)

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)

	list := bst.Postorder(bst.Root())
	exp := []int{1, 4, 6, 3, 10, 8}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("bst.Postorder(bst.Root()) = %v, want %v", list, exp)
	}
}
