package tree

import "fmt"

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

func InOrderPrint(t Tree) {
	printNode(t.Root)
}

func printNode(n *Node) {
	if n.Left != nil {
		printNode(n.Left)
	}
	fmt.Printf("%d, ", n.Value)
	if n.Right != nil {
		printNode(n.Right)
	}
}
