package main

import (
	"fmt"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/bstree"
)

func main() {
	bst := bstree.New(8, sture.Compare[int])

	bst.Insert(bst.Root(), 3)
	bst.Insert(bst.Root(), 10)
	bst.Insert(bst.Root(), 1)
	bst.Insert(bst.Root(), 6)
	bst.Insert(bst.Root(), 4)
	bst.Insert(bst.Root(), 7)

	fmt.Printf("%+v\n", bst.Search(bst.Root(), 3))
	bst.Remove(bst.Root(), 3)
	fmt.Printf("%+v\n", bst.Search(bst.Root(), 3)) // nil

	fmt.Println(bst.Inorder(bst.Root()))   // [1 4 6 7 8 10]
	fmt.Println(bst.Preorder(bst.Root()))  // [8 4 1 6 7 10]
	fmt.Println(bst.Postorder(bst.Root())) // [1 7 6 4 10 8]
}
