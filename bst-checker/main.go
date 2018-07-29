package main

import (
	"fmt"
	"interview/bst/tree"
)

func main() {
	nums := []int{1, 432, 67, 342, 5, 763, 811}
	t := tree.Tree{
		Root:   &tree.Node{Value: 300},
		Height: 1,
	}
	for _, num := range nums {
		t.Insert(num)
	}

	// t.Root.Left.Right.Left.Right = &tree.Node{
	// 	Value: 88,
	// }
	_, _, valid := checkBST(t.Root)

	println(valid)
}

func checkBST(n *tree.Node) (int, int, bool) {
	newSmallest := n.Value
	newLargest := n.Value

	if n.Left != nil {
		smallestLeft, largestLeft, valid := checkBST(n.Left)
		if !valid || largestLeft > n.Value {
			return 0, 0, false
		}
		newSmallest = smallestLeft
	}

	if n.Right != nil {
		smallestRight, largestRight, valid := checkBST(n.Right)
		if !valid || smallestRight < n.Value {
			return 0, 0, false
		}
		newLargest = largestRight
	}

	fmt.Printf("at node %d, smallest: %d, largest: %d\n", n.Value, newSmallest, newLargest)
	return newSmallest, newLargest, true
}
