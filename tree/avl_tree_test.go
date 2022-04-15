package tree_test

import (
	"reflect"
	"testing"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/tree"
)

func TestAVLTree_Insert(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	/*
			From:
						__15__
					  /		  \
					6		   23
				  /   \			 \
				4	   7		  71
			     \				 /
				  5			   50
			To:
				   __6__
				 /      \
				4		15
		         \     /  \
				  5   7	   50
		                  /  \
				  		23    71
	*/

	if avl.Root().Value() != 6 {
		t.Errorf("avl.Root().Value() = %v, want %v", avl.Root().Value(), 6)
	}

	rootLeftChild := avl.Root().Left()
	rootRightChild := avl.Root().Right()

	if rootLeftChild.Value() != 4 {
		t.Errorf("avl.Root().Left().Value() = %v, want %v", avl.Root().Left().Value(), 4)
	}

	if rootLeftChild.Right().Value() != 5 {
		t.Errorf("avl.Root().Left().Right().Value() = %v, want %v", avl.Root().Right().Left().Value(), 5)
	}

	if rootRightChild.Value() != 15 {
		t.Errorf("avl.Root().Right().Value() = %v, want %v", avl.Root().Right().Value(), 15)
	}

	if rootRightChild.Left().Value() != 7 {
		t.Errorf("avl.Root().Right().Left().Value() = %v, want %v", avl.Root().Right().Left().Value(), 7)
	}

	rightRightChild := rootRightChild.Right()
	if rightRightChild.Value() != 50 {
		t.Errorf("avl.Root().Right().Right().Value() = %v, want %v", avl.Root().Right().Right().Value(), 50)
	}

	if rightRightChild.Left().Value() != 23 {
		t.Errorf("avl.Root().Right().Right().Left().Value() = %v, want %v", avl.Root().Right().Right().Left().Value(), 23)
	}

	if rightRightChild.Right().Value() != 71 {
		t.Errorf("avl.Root().Right().Right().Right().Value() = %v, want %v", avl.Root().Right().Right().Right().Value(), 71)
	}
}

func TestAVLTree_Search(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	if avl.Search(71).Value() != 71 {
		t.Errorf("avl.Search(71).Value() = %v, want %v", avl.Search(71).Value(), 71)
	}

	if avl.Search(23).Value() != 23 {
		t.Errorf("avl.Search(23).Value() = %v, want %v", avl.Search(23).Value(), 23)
	}

	if avl.Search(15).Value() != 15 {
		t.Errorf("avl.Search(15).Value() = %v, want %v", avl.Search(15).Value(), 15)
	}

	if avl.Search(4).Value() != 4 {
		t.Errorf("avl.Search(4).Value() = %v, want %v", avl.Search(4).Value(), 4)
	}

	if avl.Search(5).Value() != 5 {
		t.Errorf("avl.Search(5).Value() = %v, want %v", avl.Search(5).Value(), 5)
	}

	if avl.Search(7).Value() != 7 {
		t.Errorf("avl.Search(7).Value() = %v, want %v", avl.Search(7).Value(), 7)
	}

	if avl.Search(50).Value() != 50 {
		t.Errorf("avl.Search(50).Value() = %v, want %v", avl.Search(50).Value(), 50)
	}

	if avl.Search(6).Value() != 6 {
		t.Errorf("avl.Search(6).Value() = %v, want %v", avl.Search(6).Value(), 6)
	}
}

func TestAVLTree_Remove(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	avl.Remove(15)
	avl.Remove(7)

	/*
			From:
				   __6__
				 /      \
				4		15
		         \     /  \
				  5   7	   50
		                  /  \
				  		23    71
			To:
				   __6__
				 /      \
				4		50
		         \     /  \
				  5   23   71
	*/

	if avl.Root().Value() != 6 {
		t.Errorf("avl.Root().Value() = %v, want %v", avl.Root().Value(), 6)
	}

	rootLeftChild := avl.Root().Left()
	rootRightChild := avl.Root().Right()

	if rootLeftChild.Value() != 4 {
		t.Errorf("avl.Root().Left().Value() = %v, want %v", avl.Root().Left().Value(), 4)
	}

	if rootLeftChild.Right().Value() != 5 {
		t.Errorf("avl.Root().Left().Right().Value() = %v, want %v", avl.Root().Right().Left().Value(), 5)
	}

	if rootRightChild.Value() != 50 {
		t.Errorf("avl.Root().Right().Value() = %v, want %v", avl.Root().Right().Value(), 50)
	}

	if rootRightChild.Left().Value() != 23 {
		t.Errorf("avl.Root().Right().Left().Value() = %v, want %v", avl.Root().Right().Left().Value(), 23)
	}

	if rootRightChild.Right().Value() != 71 {
		t.Errorf("avl.Root().Right().Right().Value() = %v, want %v", avl.Root().Right().Right().Value(), 71)
	}

	avl.Remove(71)
	avl.Remove(23)
	avl.Remove(50)

	/*
			From:
				   __6__
				 /      \
				4		50
		         \     /  \
				  5   23   71
			To:
				   5
				 /  \
				4    6
	*/

	if avl.Root().Value() != 5 {
		t.Errorf("avl.Root().Value() = %v, want %v", avl.Root().Value(), 5)
	}

	rootLeftChild = avl.Root().Left()
	rootRightChild = avl.Root().Right()

	if rootLeftChild.Value() != 4 {
		t.Errorf("avl.Root().Left().Value() = %v, want %v", avl.Root().Left().Value(), 4)
	}

	if rootRightChild.Value() != 6 {
		t.Errorf("avl.Root().Right().Value() = %v, want %v", avl.Root().Right().Value(), 6)
	}
}

func TestAVLTree_Inorder(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	list := avl.Inorder(avl.Root())
	exp := []int{4, 5, 6, 7, 15, 23, 50, 71}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("avl.Inorder(avl.Root()) = %v, want %v", list, exp)
	}
}

func TestAVLTree_Preorder(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	list := avl.Preorder(avl.Root())
	exp := []int{6, 4, 5, 15, 7, 50, 23, 71}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("avl.Preorder(avl.Root()) = %v, want %v", list, exp)
	}
}

func TestAVLTree_Postorder(t *testing.T) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	inputs := []int{15, 6, 4, 5, 7, 23, 71, 50}
	for _, n := range inputs {
		avl.Insert(n, n)
	}

	list := avl.Postorder(avl.Root())
	exp := []int{5, 4, 7, 23, 71, 50, 15, 6}

	if !reflect.DeepEqual(list, exp) {
		t.Errorf("avl.Postorder(avl.Root()) = %v, want %v", list, exp)
	}
}
