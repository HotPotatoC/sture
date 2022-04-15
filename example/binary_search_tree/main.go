package main

import (
	"fmt"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/tree"
)

type user struct {
	id    int
	name  string
	email string
}

func main() {
	bst := tree.NewBinarySearchTree[string, user](sture.Compare[string])

	inputs := []struct {
		key  string
		data user
	}{
		{"user::3", user{3, "Jane Doe", "janedoe@gmail.com"}},
		{"user::2", user{2, "Richard Roe", "richardroe@gmail.com"}},
		{"user::5", user{5, "Juan Christian", "juandotulung@gmail.com"}},
		{"user::1", user{1, "John Doe", "johndoe@gmail.com"}},
		{"user::4", user{4, "Johnny Doe", "johnnydoe@gmail.com"}},
		{"user::6", user{6, "Janie Doe", "janiedoe@gmail.com"}},
	}

	for _, inp := range inputs {
		bst.Insert(inp.key, inp.data)
	}

	node := bst.Search("user::3")
	fmt.Println(node)

	inorder := bst.Inorder()
	preorder := bst.Preorder()
	postorder := bst.Postorder()

	for _, node := range inorder {
		fmt.Printf("key: %s, name: %s\n", node.Key(), node.Value().name)
	}

	fmt.Println()

	for _, node := range preorder {
		fmt.Printf("key: %s, name: %s\n", node.Key(), node.Value().name)
	}

	fmt.Println()

	for _, node := range postorder {
		fmt.Printf("key: %s, name: %s\n", node.Key(), node.Value().name)
	}
}
