package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) insert(value, height int) int {
	if value == n.Value {
		return height
	}

	height++
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{
				Value: value,
			}
			return height
		}
		return n.Left.insert(value, height)
	}

	if n.Right == nil {
		n.Right = &Node{
			Value: value,
		}
		return height
	}
	return n.Right.insert(value, height)
}

type Tree struct {
	Root   *Node
	Height int
}

func (t *Tree) Insert(value int) {
	height := t.Root.insert(value, 0)
	if height > t.Height {
		t.Height = height
	}
}

func LevelOrder(tree Tree) [][]int {
	return levelOrder(tree.Root)
}

func levelOrder(node *Node) [][]int {
	// base case return the current node as a single level
	if node.Left == nil && node.Right == nil {
		return [][]int{
			{node.Value},
		}
	}

	// Get arrays for sub nodes
	var leftLevels [][]int
	var rightLevels [][]int
	if node.Left != nil {
		leftLevels = levelOrder(node.Left)
	}
	if node.Right != nil {
		rightLevels = levelOrder(node.Right)
	}

	// Merge lists of subnodes
	var i int
	for i = 0; i < len(leftLevels); i++ {
		if i < len(rightLevels) {
			leftLevels[i] = append(leftLevels[i], rightLevels[i]...)
		}
	}
	if i < len(rightLevels) {
		leftLevels = append(leftLevels, rightLevels[i:]...)
	}

	// Prepend current nodes value as a single element list at the top.
	leftLevels = append([][]int{{node.Value}}, leftLevels...)
	return leftLevels
}

func InOrderPrint(t Tree) {
	printNodeInOrder(t.Root)
}

func printNodeInOrder(n *Node) {
	if n.Left != nil {
		printNodeInOrder(n.Left)
	}
	fmt.Printf("%d, ", n.Value)
	if n.Right != nil {
		printNodeInOrder(n.Right)
	}
}

func PostOrderPring(t Tree) {
	printNodePostOrder(t.Root)
}

func printNodePostOrder(n *Node) {
	if n.Right != nil {
		printNodePostOrder(n.Right)
	}
	fmt.Printf("%d, ", n.Value)
	if n.Left != nil {
		printNodePostOrder(n.Left)
	}
}
