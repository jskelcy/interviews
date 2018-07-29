package main

import (
	"interview/bst/tree"
)

func main() {
	nums := []int{1, 432, 67, 342, 5, 763, 423, 6, 534, 675, 53, 24, 55, 6, 6787, 23, 65, 487, 56, 875, 43, 545, 7, 46, 67, 875, 657}
	t := tree.Tree{
		Root:   &tree.Node{Value: 1},
		Height: 1,
	}
	for _, num := range nums {
		t.Insert(num)
	}

	tree.InOrderPrint(t)
}
